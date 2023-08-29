package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AsteroidBoldDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="currentColor"><path d="M12 2C6.477 2 2 6.477 2 12c0 .445.029.883.085 1.312a6 6 0 0 1 7.297 8.342c.835.226 1.712.346 2.618.346c4.879 0 8.941-3.494 9.823-8.116a6.002 6.002 0 0 1-3.505-9.636A9.959 9.959 0 0 0 12 2Z" opacity=".5"/><path d="M2.085 13.312a10.01 10.01 0 0 0 7.297 8.342a6 6 0 0 0-7.297-8.342Zm19.738.573c.116-.61.177-1.24.177-1.885a9.98 9.98 0 0 0-3.682-7.752a6.002 6.002 0 0 0 3.505 9.637ZM16 16a1 1 0 1 1-2 0a1 1 0 0 1 2 0Zm-3-7.5a2.5 2.5 0 1 1-5 0a2.5 2.5 0 0 1 5 0Z"/></g>`),
		g.Group(children),
	)
}