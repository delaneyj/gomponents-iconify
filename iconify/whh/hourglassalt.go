package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Hourglassalt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M960.356 128h-64v64q0 99-46.5 183t-125.5 137q79 53 125.5 137.5t46.5 182.5v64h64q26 0 45 18.5t19 45.5t-19 45.5t-45 18.5h-896q-26 0-45-18.5t-19-45t19-45.5t45-19h64v-64q0-98 46.5-182.5t125.5-137.5q-79-53-125.5-137t-46.5-183v-64h-64q-26 0-45-18.5T.356 64t19-45.5t45-18.5h896q26 0 45 18.5t19 45.5t-19 45.5t-45 18.5zm-384 456v192l192 120v-64q0-88-54-157t-138-91zm-320 248v64l192-120V584q-84 22-138 91t-54 157zm512-704h-512v64q0 68 35 128h442q35-60 35-128v-64z"/>`),
		g.Group(children),
	)
}