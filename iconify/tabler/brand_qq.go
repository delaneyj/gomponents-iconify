package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BrandQq(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M6 9.748a14.716 14.716 0 0 0 11.995-.052C18.27.46 6.891-1.56 6 9.748zM18 10c.984 2.762 1.949 4.765 2 7.153c.014.688-.664 1.346-1.184.303C18.47 16.76 17.864 16.275 17 16m0 0c.031 1.831.147 3.102-1 4m-8 0c-1.099-.87-.914-2.24-1-4m-1-6c-.783 2.338-1.742 4.12-1.968 6.43c-.217 2.227.716 1.644 1.16.917C5.488 16.86 6.09 16.413 7 16m8.898-3l-.476-2M8 20l-1.5 1c-.5.5-.5 1 .5 1h10c1 0 1-.5.5-1L16 20"/><path d="M12.75 7a1 1 0 1 0 2 0a1 1 0 1 0-2 0m-3.5 0a1 1 0 1 0 2 0a1 1 0 1 0-2 0"/></g>`),
		g.Group(children),
	)
}