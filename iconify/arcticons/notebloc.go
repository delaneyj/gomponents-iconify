package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Notebloc(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m5.897 35.48l13.953-.923M4.603 15.931l13.954-.923m-14.051-.57a2.223 2.223 0 0 1 2.092-2.364l9.473-.626a2.227 2.227 0 0 1 2.386 2.067l1.48 22.379a2.228 2.228 0 0 1-2.092 2.364l-9.474.626a2.22 2.22 0 0 1-2.385-2.067c-.384-7.467-.946-14.922-1.48-22.38Zm14.171 1.694l1.66-9.463l20.215 2.978l-3.714 27.781l-16.988-2.871m19.896-18.943l3.754 1.95l-5.285 10.11M33.599 37.1l-2.206 4.232l-11.425-5.666m2.39-24.86l14.401 2.056m-14.73.92l13.689 1.893m-14.93 6.421l13.965 1.973m-13.53-4.799l14.4 2.055m-13.94-4.8l11.916 1.604m-13.235 6.774l10.451 1.509m-11.38 1.168l14.286 2.11"/><path fill="currentColor" d="M13.738 36.72a.75.75 0 1 1-.798-.698a.751.751 0 0 1 .798.698Z"/>`),
		g.Group(children),
	)
}