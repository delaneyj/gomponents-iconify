package fxemoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Heavyexclaimationmarksymbol(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="#2B3B47" d="M305.486 405.211c0 27.191-21.752 49.486-49.486 49.486c-27.734 0-49.486-22.295-49.486-49.486c0-27.19 21.752-49.485 49.486-49.485c27.734 0 49.486 22.296 49.486 49.485zm-26.808-98.971h-45.873c-6.8 0-12.398-5.347-12.708-12.14l-9.711-212.197c-.332-7.247 5.454-13.303 12.708-13.303h65.781c7.266 0 13.056 6.075 12.707 13.332l-10.197 212.197c-.325 6.781-5.918 12.111-12.707 12.111z"/>`),
		g.Group(children),
	)
}