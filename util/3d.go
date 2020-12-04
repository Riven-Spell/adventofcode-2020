package util

type Vector3D struct {
	X, Y, Z int64
}

func (p Vector3D) Add(p2 Vector3D) Vector3D {
	return Vector3D{
		X: p.X + p2.X,
		Y: p.Y + p2.Y,
		Z: p.Z + p2.Z,
	}
}