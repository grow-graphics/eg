package main

import (
	"math/rand"

	"graphics.gd/classdb/AnimatedSprite2D"
	"graphics.gd/classdb/RigidBody2D"
	"graphics.gd/classdb/SpriteFrames"
	"graphics.gd/variant/StringName"
)

type Mob struct {
	RigidBody2D.Extension[Mob] `gd:"DodgeTheCreepsMob"`

	AnimatedSprite2D AnimatedSprite2D.Instance
}

func (m *Mob) Ready() {
	mobTypes := SpriteFrames.Instance(m.AnimatedSprite2D.SpriteFrames()).GetAnimationNames()
	pick := int64(rand.Int() % len(mobTypes))
	AnimatedSprite2D.Advanced(m.AnimatedSprite2D).Play(StringName.New(mobTypes[pick]), 1, false)
}

func (m *Mob) OnVisibleOnScreenNotifier2DScreenExited() { m.AsNode().QueueFree() }
