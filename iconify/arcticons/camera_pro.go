package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CameraPro(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m19.83 8.77l-2.77 2.84H6.29A1.79 1.79 0 0 0 4.5 13.4v23.22c0 .99.8 1.794 1.79 1.8h35.42a1.8 1.8 0 0 0 1.79-1.8V13.4a1.79 1.79 0 0 0-1.79-1.79H30.94l-2.77-2.84h-8.34Zm18.916 5.58c1.165 0 2.109.895 2.109 2s-.944 2-2.109 2c-1.158-.113-2-1.106-1.88-2.205c.103-.942.888-1.687 1.88-1.795Z"/><g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><circle cx="24" cy="26.223" r="8.507"/><path d="M26.454 21.972h-4.908l-2.455 4.251l2.455 4.251h4.908l2.455-4.251l-2.455-4.251zm2.455 4.251l2.46 4.262m-14.732-8.513l2.454 4.251M24 17.716l-2.454 4.256m4.908 8.502L24 34.73m-2.454-4.256H16.63m9.824-8.502h4.915"/></g>`),
		g.Group(children),
	)
}