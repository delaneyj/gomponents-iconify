package pajamas

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MarqueeSelection(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M1.75 1H1v2.25a.75.75 0 0 0 1.5 0V2.5h.75a.75.75 0 0 0 0-1.5h-1.5Zm0 5a.75.75 0 0 1 .75.75v2.5a.75.75 0 0 1-1.5 0v-2.5A.75.75 0 0 1 1.75 6Zm7.106 2.856l1.507 3.517l2.01-2.01l-3.517-1.507Zm-.558-1.871l5.344 2.29a1 1 0 0 1 .314 1.627l-.997.996l1.821 1.822a.75.75 0 1 1-1.06 1.06l-1.822-1.821l-.996.997a1 1 0 0 1-1.627-.314l-2.29-5.344c-.356-.83.483-1.669 1.313-1.313ZM1.75 15H1v-2.25a.75.75 0 0 1 1.5 0v.75h.75a.75.75 0 0 1 0 1.5h-1.5ZM6 1.75A.75.75 0 0 1 6.75 1h2.5a.75.75 0 0 1 0 1.5h-2.5A.75.75 0 0 1 6 1.75ZM14.25 1H15v2.25a.75.75 0 0 1-1.5 0V2.5h-.75a.75.75 0 0 1 0-1.5h1.5Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}