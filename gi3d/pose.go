// Copyright (c) 2019, The GoKi Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gi3d

import "github.com/goki/gi/mat32"

// Pose contains the full specification of a given object's position and orientation
type Pose struct {
	Pos         mat32.Vec3 `desc:"position of center of object"`
	Scale       mat32.Vec3 `desc:"scale (relative to parent)"`
	Quat        mat32.Quat `desc:"Node rotation specified as a Quat (relative to parent)"`
	Matrix      mat32.Mat4 `desc:"Local matrix. Contains all position/rotation/scale information (relative to parent)"`
	ParMatrix   mat32.Mat4 `desc:"Parent's world matrix -- we cache this so that we can independently update our own matrix"`
	WorldMatrix mat32.Mat4 `desc:"World matrix. Contains all absolute position/rotation/scale information (i.e. relative to very top parent, generally the scene)"`
	MVMatrix    mat32.Mat4 `desc:"model * view matrix -- tranforms into camera-centered coords"`
	MVPMatrix   mat32.Mat4 `desc:"model * view * projection matrix -- full final render matrix"`
	NormMatrix  mat32.Mat3 `desc:"normal matrix based on MVMatrix"`
}

func (ps *Pose) Defaults() {
	ps.Scale.Set(1, 1, 1)
	ps.Quat.SetIdentity()
}

// UpdateMatrix updates the local transform matrix based on its position, quaternion, and scale.
// Also checks for degenerate nil values
func (ps *Pose) UpdateMatrix() {
	if ps.Scale.IsNil() {
		ps.Scale.Set(1, 1, 1)
	}
	if ps.Quat.IsNil() {
		ps.Quat.SetIdentity()
	}
	ps.Matrix.SetTransform(ps.Pos, ps.Quat, ps.Scale)
}

// UpdateWorldMatrix updates the world transform matrix based on Matrix and parent's WorldMatrix.
// Also calls UpdateMatrix
func (ps *Pose) UpdateWorldMatrix(parWorld *mat32.Mat4) {
	ps.UpdateMatrix()
	ps.ParMatrix = *parWorld
	ps.WorldMatrix.MulMatrices(parWorld, &ps.Matrix)
}

// UpdateMVPMatrix updates the model * view, * projection matricies based on camera view, prjn matricies
// Assumes that WorldMatrix has been updated
func (ps *Pose) UpdateMVPMatrix(viewMat, prjnMat *mat32.Mat4) {
	ps.MVMatrix.MulMatrices(viewMat, &ps.WorldMatrix)
	ps.NormMatrix.SetNormalMatrix(&ps.MVMatrix)
	ps.MVPMatrix.MulMatrices(prjnMat, &ps.MVMatrix)
}

// MoveOnAxis moves (translates) the specified distance on the specified local axis.
func (ps *Pose) MoveOnAxis(x, y, z, dist float32) {
	ps.Pos.SetAdd(mat32.NewVec3(x, y, z).MulQuat(ps.Quat).MulScalar(dist))
}

// SetEulerRotation sets the rotation in Euler angles (radians).
func (ps *Pose) SetEulerRotation(x, y, z float32) {
	rot := mat32.NewVec3(x, y, z)
	ps.Quat.SetFromEuler(rot)
}

// EulerRotation returns the current rotation in Euler angles (radians).
func (ps *Pose) EulerRotation() mat32.Vec3 {
	rot := mat32.Vec3{}
	rot.SetEulerAnglesFromQuat(ps.Quat)
	return rot
}

// SetAxisRotation sets rotation from local axis and angle in radians.
func (ps *Pose) SetAxisRotation(x, y, z, angle float32) {
	axis := mat32.NewVec3(x, y, z)
	ps.Quat.SetFromAxisAngle(axis, angle)
}

// RotateOnAxis rotates around the specified local axis the specified angle in radians.
func (ps *Pose) RotateOnAxis(x, y, z, angle float32) {
	axis := mat32.NewVec3(x, y, z)
	rotQuat := mat32.Quat{}
	rotQuat.SetFromAxisAngle(axis, angle)
	ps.Quat.Mul(rotQuat)
}

// SetMatrix sets the local transformation matrix and updates Pos, Scale, Quat.
func (ps *Pose) SetMatrix(m *mat32.Mat4) {
	ps.Matrix = *m
	ps.Pos, ps.Quat, ps.Scale = ps.Matrix.Decompose()
}

// WorldPos returns the current world position.
func (ps *Pose) WorldPos() mat32.Vec3 {
	pos := mat32.Vec3{}
	pos.SetFromMatrixPos(&ps.WorldMatrix)
	return pos
}

// WorldQuat returns the current world quaternion.
func (ps *Pose) WorldQuat() mat32.Quat {
	_, quat, _ := ps.WorldMatrix.Decompose()
	return quat
}

// WorldEulerRotation returns the current world rotation in Euler angles.
func (ps *Pose) WorldEulerRotation() mat32.Vec3 {
	ang := mat32.Vec3{}
	ang.SetEulerAnglesFromQuat(ps.Quat)
	return ang
}

// WorldScale returns he current world scale.
func (ps *Pose) WorldScale() mat32.Vec3 {
	_, _, scale := ps.WorldMatrix.Decompose()
	return scale
}