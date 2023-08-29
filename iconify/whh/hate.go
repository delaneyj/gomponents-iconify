package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Hate(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M736.528 576q-4 0-10.5.5t-23 4t-29 9.5t-23 19t-10.5 31q0 21 10 39.5t22 32.5t22 45.5t10 74.5v128q0 27-18.5 45.5t-45.5 18.5t-45.5-18.5t-18.5-45.5v-32q0-33-21.5-67.5t-49.5-60t-64.5-62t-56.5-66.5q-7-11-21-27t-44-40.5t-64.5-44t-85.5-34t-105-14.5q-3 0-4-1q-23-4-41.5-48t-18.5-79V128q0-31 18-47.5t46-16.5q32 0 128-32t128-32h512q26 0 45 18.5t19 45t-19 45.5t-45 19h64q26 0 45 18.5t19 45t-19 45.5t-45 19q26 0 45 18.5t19 45t-18.5 45.5t-45.5 19h64q27 0 45.5 18.5t18.5 45.5q0 49-85.5 88.5t-202.5 39.5z"/>`),
		g.Group(children),
	)
}