package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EnglishMustache(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linejoin="round" stroke-width="4"><path d="M10.064 24.058c3.287-1.073 3.465-5.092 7.797-5.94c3.465-.68 5.547 1.708 6.064 4.243c.433 2.122-.866 7.639-7.797 7.639C7.465 30 4.578 24.483 4 22.361c.866.849 3.465 2.546 6.064 1.697Z"/><path d="M37.936 24.058c-3.288-1.073-3.465-5.092-7.797-5.94c-3.465-.68-5.547 1.708-6.064 4.243c-.433 2.122.866 7.639 7.797 7.639c8.663 0 11.55-5.517 12.128-7.639c-.866.849-3.465 2.546-6.064 1.697Z"/></g>`),
		g.Group(children),
	)
}