package openmoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Dizzy(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 72 72"),
		g.Raw(`<path fill="#fcea2b" d="M14.827 18.535c-.336-.226 80.195-10.75 29.962 39.444L32.69 48.48s47.75-18.573-17.863-29.945z"/><path fill="#f1b31c" d="m31.298 48.361l3.791 3.112c11.064-4.99 39.602-20.574-.447-33.844c-10.641-.467-20.642.906-20.642.906c64.003 11.049 17.298 29.826 17.298 29.826z"/><path fill="#fcea2b" d="m35 30.403l4.944 10.018L51 42.028l-8 7.798l1.889 11.011L35 55.638l-9.889 5.199L27 49.826l-8-7.798l11.056-1.607z"/><g fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-miterlimit="10" stroke-width="2"><path d="M48.665 34.784c.72-5.455-6.509-11.55-33.838-16.249c0 0 72.994-9.998 34.807 33.068"/><path d="m35 30.403l4.944 10.018L51 42.028l-7 7.798l1.889 11.011L35 55.638l-9.889 5.199L27 49.826l-8-7.798l11.056-1.607z"/></g>`),
		g.Group(children),
	)
}