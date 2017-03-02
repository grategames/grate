package grate

import "time"

var start time.Time = time.Now()


type Animate struct {}

//Return a Scale animation which scales from start to end over the sepcified time.
func (Animate) Scale(start, end, time float64) Animation {
	return Animation{Scale: AnimationProperty{ Start:start, End:end, Time:time, Active: true}}
}

//Return a Translate animation which translate from start to end over the sepcified time.
func (Animate) Translate(start, end, time float64) Animation {
	return Animation{Translate: AnimationProperty{ Start:start, End:end, Time:time, Active: true}}
}

//Return a Rotate animation which rotate from start to end over the sepcified time.
func (Animate) Rotate(start, end, time float64) Animation {
	return Animation{Rotate: AnimationProperty{ Start:start, End:end, Time:time, Active: true}}
}

type AnimationProperty struct {
	Active bool
	Start, End, Time float64
}

func (prop *AnimationProperty) Reverse() {
	prop.Start, prop.End = prop.End, prop.Start
}


//An animation is a translation, scale or rotation over time.
//They can be applied to an image.
type Animation struct {
	Started bool
	Finished bool
	StartTime float64

	Scale AnimationProperty
	Translate AnimationProperty
	Rotate AnimationProperty
}

//Update the animation, this will return true when the animation is complete.
func (anim *Animation) Update() bool {
	if !anim.Started {
		anim.Started = true
		anim.StartTime = time.Since(start).Seconds()
	}
	if anim.Finished {
		return true
	}
	return false
}

// Precise method, which guarantees v = v1 when t = 1.
func lerp(v0, v1, t float64) float64 {
  return (1 - t) * v0 + t * v1;
}

//Reverse an animation.
func (anim *Animation) Reverse() {
	anim.StartTime = time.Since(start).Seconds()
	anim.Finished = false
	anim.Scale.Reverse()
	anim.Translate.Reverse()
	anim.Rotate.Reverse()
}

//Apply an animation on an image.
func (anim *Animation) Apply(img Image) {

	if anim.Scale.Active {
		if time.Since(start).Seconds()-anim.StartTime > anim.Scale.Time {
			img.Scale(anim.Scale.End, anim.Scale.End)
			anim.Finished = true
		} else {
			img.Scale(lerp(anim.Scale.Start, anim.Scale.End, (time.Since(start).Seconds()-anim.StartTime)/anim.Scale.Time), 
				lerp(anim.Scale.Start, anim.Scale.End, (time.Since(start).Seconds()-anim.StartTime)/anim.Scale.Time))
		}
	}
	
	if anim.Translate.Active {
		if time.Since(start).Seconds()-anim.StartTime > anim.Translate.Time {
			img.Translate(anim.Translate.End, anim.Translate.End)
			anim.Finished = true
		} else {
			img.Translate(lerp(anim.Translate.Start, anim.Translate.End, (time.Since(start).Seconds()-anim.StartTime)/anim.Translate.Time), 
				lerp(anim.Translate.Start, anim.Translate.End, (time.Since(start).Seconds()-anim.StartTime)/anim.Translate.Time))
		}
	}
	
	if anim.Rotate.Active {
		if time.Since(start).Seconds()-anim.StartTime > anim.Rotate.Time {
			img.Rotate(anim.Rotate.End)
			anim.Finished = true
		} else {
			img.Rotate(lerp(anim.Rotate.Start, anim.Rotate.End, (time.Since(start).Seconds()-anim.StartTime)/anim.Rotate.Time))
		}
	}
}
