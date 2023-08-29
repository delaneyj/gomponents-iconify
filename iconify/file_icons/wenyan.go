package file_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Wenyan(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M0 113V85c36.35 0 42-17.16 42-23c0-11.797-20.708-14.138-20-32c1.6-40.356 58.163-39.615 68 0c11.493 46.284-21.558 87.005-90 83zm125-7h175V65H125v41zM0 187h125v-41H0v41zm173 0h127v-41H173v41zM0 267h125v-41H0v41zm173 0h127v-41H173v41zm-48 81v-41H20v205h105v-41H66V348h59zm48-41v41h59v123h-59v41h105V307H173z"/>`),
		g.Group(children),
	)
}