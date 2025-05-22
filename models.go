package lucide

import "github.com/a-h/templ"

// Props defines the properties that can be set for an icon.
type Props struct {
	ID          string
	Size        int
	Color       string
	Fill        string
	Stroke      string
	StrokeWidth float32
	Class       string           // additional css classes
	Attrs       templ.Attributes // extra attributes
}

func (props *Props) size() int {
	if props.Size > 0 {
		return props.Size
	}
	return 24
}

func (props *Props) color() string {
	if props.Color != "" {
		return props.Color
	}
	return "currentColor"
}

func (props *Props) strokeColor() string {
	return props.color()
}

func (props *Props) strokeWidth() float32 {
	if props.StrokeWidth > 0 {
		return props.StrokeWidth
	}
	return 2
}
