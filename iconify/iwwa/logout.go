package iwwa

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Logout(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 40 40"),
		g.Raw(`<path fill="currentColor" d="M28.846 33.405H5.596a1.782 1.782 0 0 1-1.781-1.781V8.375c0-.982.799-1.781 1.781-1.781h23.249c.982 0 1.781.798 1.781 1.781V14a.46.46 0 1 1-.922 0V8.375a.86.86 0 0 0-.858-.858H5.596a.86.86 0 0 0-.858.858v23.25a.86.86 0 0 0 .858.858h23.249a.86.86 0 0 0 .858-.858v-5.568a.46.46 0 1 1 .922 0v5.568a1.78 1.78 0 0 1-1.779 1.78zM33.2 27.02a.46.46 0 0 1-.333-.781l6.03-6.253l-6.005-6.225a.46.46 0 1 1 .665-.64l6.313 6.545a.46.46 0 0 1 .001.64l-6.339 6.573a.46.46 0 0 1-.332.141z"/><path fill="currentColor" d="M39.538 20.447H17.22a.46.46 0 1 1 0-.922h22.318a.46.46 0 1 1 0 .922z"/>`),
		g.Group(children),
	)
}