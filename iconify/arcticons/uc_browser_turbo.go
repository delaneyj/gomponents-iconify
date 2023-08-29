package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func UcBrowserTurbo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><circle cx="19.584" cy="35.252" r="7.248"/><path d="M29.384 38.58s5.985-1.098 5.985 3.92H16.524c-6.184 0-11.197-5.013-11.197-11.197c0-2.9 1.1-5.548 2.914-7.533c3.51-3.84 5.1-5.194 5.1-8.375s-3.953-5.124-8.014-2.65C8.806 6.678 11.574 5.5 17.228 5.5s8.305 4.712 8.305 8.835c0 8.187-13.197 8.893-13.197 20.917"/><path d="M29.384 38.58a9.131 9.131 0 0 0-15.58-9.521"/><path d="M18.141 25.52c5.272-3.705 12.575-2.233 17.228 2.185c5.36-.943 7.304 2.356 7.304 2.356c-2.28-.178-5.034.656-6.854 1.343a3.04 3.04 0 0 1-3.336-.79c-3.593-3.961-9.133-7.585-14.342-5.093Z"/><path d="M22.98 23.573s2.789-2.863 8.973-6.176c-.236-3.003-.088-4.13 1.414-5.065c2.65.67 3.18 3.829 3.18 3.829c4.889 1.649 6.832 8.304 5.007 9.247s-7.741.256-11.353-.98"/></g><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M19.44 31.252V38.1c0 .635.516 1.15 1.152 1.15h.345m-2.705-6.099h2.417"/>`),
		g.Group(children),
	)
}