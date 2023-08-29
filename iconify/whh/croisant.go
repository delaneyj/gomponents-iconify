package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Croisant(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M1023 288q0 8-1 14t-5 9.5t-6.5 5.5t-10 2.5t-11.5.5h-30q-12 0-27-10t-27-21.5t-31-21.5t-40-11l43-154q45 31 76 62l49 49q8 8 13 15t6.5 16t2 13t0 15.5t-.5 15.5zm-319-32h-67q-31-54-83-123.5T455 10q189-35 361 54l-55 192h-57zM552 371L371 552q-19 18-26 21t-32 3q-31 0-131-72T27 377q-17-17-24-47t2-56q35-92 106-163.5T274 5q26-9 56-2t47 23q55 56 127 156t72 131q0 25-3 32t-21 26zM256 834q1 21 11 40t21.5 31t21.5 27t10 27q0 33-6.5 48.5T288 1023q-4 0-15.5.5t-15.5 0t-13-2t-16-6.5t-15-13l-49-49q-31-31-62-75zM64 816q-89-172-54-361q53 47 122.5 99T256 638v123z"/>`),
		g.Group(children),
	)
}