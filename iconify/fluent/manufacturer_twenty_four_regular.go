package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ManufacturerTwentyFourRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<g fill="none"><path d="M2.924 4.129A2.25 2.25 0 0 1 5.171 2H6.33a2.25 2.25 0 0 1 2.247 2.129l.838 15.5A2.25 2.25 0 0 1 7.167 22H4.333a2.25 2.25 0 0 1-2.246-2.371l.837-15.5zM5.171 3.5a.75.75 0 0 0-.749.71l-.838 15.5a.75.75 0 0 0 .75.79h2.833a.75.75 0 0 0 .749-.79l-.838-15.5a.75.75 0 0 0-.75-.71H5.172z" fill="currentColor"/><path d="M9.512 22H19.75A2.25 2.25 0 0 0 22 19.75V6a.75.75 0 0 0-1.27-.541l-4.98 4.781V6a.75.75 0 0 0-1.228-.578L9.856 9.284l.1 1.864l4.294-3.554V12a.75.75 0 0 0 1.27.541L20.5 7.76v11.99a.75.75 0 0 1-.75.75h-1V16A1.75 1.75 0 0 0 17 14.25h-4A1.75 1.75 0 0 0 11.25 16v4.5h-.92a3.248 3.248 0 0 1-.818 1.5zm7.738-6v4.25h-4.5V16a.25.25 0 0 1 .25-.25h4a.25.25 0 0 1 .25.25z" fill="currentColor"/></g>`),
		g.Group(children),
	)
}