package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Tunnelbear(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><ellipse cx="24" cy="10.934" rx="13.143" ry="3.772"/><path d="M40.706 10.894c0 3.531-7.48 6.394-16.706 6.394S7.294 14.425 7.294 10.894m0 0C7.294 7.363 14.774 4.5 24 4.5s16.706 2.863 16.706 6.394v6.351c0 3.532-7.48 6.394-16.706 6.394S7.294 20.777 7.294 17.245v-6.35Z"/><path d="M39.196 20.157v17.487c0 3.234-6.803 5.856-15.196 5.856S8.804 40.878 8.804 37.644V20.157"/></g><g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M14.665 39.329c1.231-1.362 4.722-.27 6.97-.27c1.014 0 2.128-.317 2.805-2.226c1.034-2.914 2.598-7.446 3.276-9.68"/><path d="M33.813 27.154H20.97c-6.638 0-7.678 4.3-6.038 6.311c2.99 3.668 8.104.777 6.782-2.184c-.744-1.665-2.892-1.922-3.863-.751"/></g>`),
		g.Group(children),
	)
}