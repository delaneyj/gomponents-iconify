package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AttachTextTwentyFourRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M7.25 3.5a4.25 4.25 0 0 1 4.245 4.044l.005.206V18.5a2.5 2.5 0 0 1-4.995.164L6.5 18.5V9a.75.75 0 0 1 1.493-.102L8 9v9.5a1 1 0 0 0 1.993.117L10 18.5V7.75a2.75 2.75 0 0 0-5.495-.168L4.5 7.75V17a.75.75 0 0 1-1.493.102L3 17V7.75A4.25 4.25 0 0 1 7.25 3.5Zm9 12.5a.75.75 0 0 1 .102 1.493l-.102.007h-2.5a.75.75 0 0 1-.102-1.493L13.75 16h2.5Zm4-3a.75.75 0 0 1 .102 1.493l-.102.007h-6.5a.75.75 0 0 1-.102-1.493L13.75 13h6.5Zm0-3a.75.75 0 0 1 .102 1.493l-.102.007h-6.5a.75.75 0 0 1-.102-1.493L13.75 10h6.5Zm0-3a.75.75 0 0 1 .102 1.493l-.102.007h-6.5a.75.75 0 0 1-.102-1.493L13.75 7h6.5Z"/>`),
		g.Group(children),
	)
}