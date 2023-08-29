package flagpack

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func My(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 24"),
		g.Raw(`<defs><path id="flagpackMy0" fill="#fff" d="M0 0h32v24H0z"/></defs><g fill="none"><g clip-path="url(#flagpackMy1)"><use href="#flagpackMy0"/><path fill="#F7FCFF" fill-rule="evenodd" d="M0 0h32v24H0V0Z" clip-rule="evenodd"/><path fill="#E31D1C" d="M.027 5h32v2.9h-32zm0 5.2h32v2.9h-32zm.085 5h32v2.5h-32zm0 5h32v2.7h-32z"/><path fill="#E31D1C" fill-rule="evenodd" d="M0 0h32v2.5H0V0Z" clip-rule="evenodd"/><path fill="#3D58DB" fill-rule="evenodd" d="M0 0h16v12H0V0Z" clip-rule="evenodd"/><path fill="#FECA00" fill-rule="evenodd" d="M3.594 6.202c0 1.376.648 2.678 1.966 2.678c1.98 0 2.364-.422 3.287-1.004c.219.49-.518 2.588-3.314 2.588C3.297 10.42 1.49 8.623 1.49 6.202c0-2.778 2.044-4.279 3.987-4.262c1.716 0 3.536.974 3.48 2.216c-.808-.79-1.686-.79-3.147-.79c-1.462 0-2.216 1.46-2.216 2.836Z" clip-rule="evenodd"/><path fill="#FECA00" fill-rule="evenodd" d="m11 7.3l-.927 1.353l.045-1.64l-1.545.55l1-1.3L8 5.8l1.573-.464l-1-1.3l1.545.55l-.045-1.64L11 4.3l.927-1.353l-.045 1.64l1.545-.55l-1 1.3L14 5.8l-1.573.464l1 1.3l-1.545-.55l.045 1.64L11 7.3Z" clip-rule="evenodd"/></g><defs><clipPath id="flagpackMy1"><use href="#flagpackMy0"/></clipPath></defs></g>`),
		g.Group(children),
	)
}