package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PillsBroken(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none"><path stroke="currentColor" stroke-linecap="round" stroke-width="1.5" d="M17.845 6.155s-.433 2.245-2.94 4.751c-2.505 2.506-4.75 2.938-4.75 2.938m11.628-2.806a5.413 5.413 0 0 0-1.376-2.32l-5.126-5.126a5.437 5.437 0 1 0-7.689 7.69l5.127 5.125A5.437 5.437 0 0 0 21.425 15M14.5 6.5L13 5"/><path fill="currentColor" d="M6.888 10.868a.75.75 0 1 0-.316-1.466l.316 1.466Zm-1.075 9.906a.75.75 0 1 0-.626 1.364l.626-1.364Zm-3.95-1.96a.75.75 0 1 0 1.363-.627l-1.364.626ZM2.75 16a5.252 5.252 0 0 1 4.138-5.132l-.316-1.466A6.752 6.752 0 0 0 1.25 16h1.5Zm10.389 1.08A5.252 5.252 0 0 1 8 21.25v1.5a6.75 6.75 0 0 0 6.607-5.364l-1.468-.306ZM8 21.25c-.782 0-1.522-.17-2.187-.476l-.626 1.364c.857.393 1.81.612 2.813.612v-1.5Zm-4.774-3.063A5.228 5.228 0 0 1 2.75 16h-1.5c0 1.003.219 1.956.612 2.813l1.364-.626Z"/></g>`),
		g.Group(children),
	)
}