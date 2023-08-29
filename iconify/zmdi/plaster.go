package zmdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Plaster(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 432 384"),
		g.Raw(`<path fill="currentColor" d="m336 216l86 85q6 7 6 15.5t-6 14.5l-93 93q-6 6-15 6t-15-6l-85-85l-85 85q-6 6-15 6t-15-6L6 331q-6-6-6-14.5T6 301l85-85l-85-84q-6-7-6-15.5T6 101L99 9q6-6 14.5-6T129 9l85 85l85-85q6-6 15-6t15 6l92 92q7 7 7 15.5t-7 15.5zm-122-64q-9 0-15 6.5t-6 15t6 15t15 6.5t15.5-6.5t6.5-15t-6.5-15T214 152zm-100 42l77-78l-77-77l-78 78zm57.5 43q8.5 0 15-6t6.5-15t-6.5-15t-15-6t-15 6t-6.5 15t6.5 15t15 6zm42.5 43q9 0 15.5-6.5t6.5-15t-6.5-15T214 237t-15 6.5t-6 15t6 15t15 6.5zm43-85q-9 0-15 6t-6 15t6 15t15 6t15-6t6-15t-6-15t-15-6zm57 199l77-78l-77-77l-78 78z"/>`),
		g.Group(children),
	)
}