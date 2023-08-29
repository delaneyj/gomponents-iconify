package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AirbudsLeftBoldDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="currentColor"><path d="M4.382 22C3.067 22 2 20.88 2 19.5h4.765c0 1.38-1.067 2.5-2.383 2.5Z"/><path fill-rule="evenodd" d="M11 7.889V5.542c0-.194 0-.29-.004-.372c-.08-1.713-1.385-3.082-3.017-3.166C7.902 2 7.81 2 7.625 2c-.307 0-.46 0-.59.007c-2.72.14-4.895 2.422-5.029 5.276C2 7.419 2 7.58 2 7.903v9.93h4.765v-5.5c0-.613.474-1.11 1.059-1.11C9.578 11.222 11 9.73 11 7.888ZM9.324 5.125c0-.46-.336-.833-.75-.833c-.415 0-.75.373-.75.833v2.778c0 .46.335.833.75.833c.414 0 .75-.373.75-.833V5.125Z" clip-rule="evenodd"/><path d="M16.5 22a5.5 5.5 0 1 1 0-11a5.5 5.5 0 0 1 0 11Z" opacity=".5"/><path d="M15.25 13a.75.75 0 0 1 .75.75V18h2.25a.75.75 0 0 1 0 1.5h-3a.75.75 0 0 1-.75-.75v-5a.75.75 0 0 1 .75-.75Z"/><path fill-rule="evenodd" d="M13.265 4.95a.75.75 0 0 0 .586.885a4.256 4.256 0 0 1 3.314 3.314a.75.75 0 0 0 1.47-.298a5.756 5.756 0 0 0-4.486-4.486a.75.75 0 0 0-.884.586Z" clip-rule="evenodd" opacity=".5"/></g>`),
		g.Group(children),
	)
}