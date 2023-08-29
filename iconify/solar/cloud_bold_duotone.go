package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CloudBoldDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="currentColor"><path fill-rule="evenodd" d="M22 14.353C22 17.472 19.442 20 16.286 20h-5.787a7.5 7.5 0 0 1 7.487-11.853c.079.281.136.571.17.868C20.392 9.78 22 11.881 22 14.353Z" clip-rule="evenodd" opacity=".5"/><path d="M12.476 4C9.32 4 6.762 6.528 6.762 9.647c0 .69.125 1.35.354 1.962a4.351 4.351 0 0 0-.83-.08C3.919 11.53 2 13.426 2 15.765C2 18.104 3.919 20 6.286 20h4.213a7.5 7.5 0 0 1 7.487-11.853a5.255 5.255 0 0 0-.046-.158C17.224 5.68 15.048 4 12.476 4Z"/></g>`),
		g.Group(children),
	)
}