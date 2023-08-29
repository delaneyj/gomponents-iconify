package vs

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TableAlt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1792 1792"),
		g.Raw(`<path fill="currentColor" d="M289 192h1086q202 185 267 245q22 11 22 34v141q0 22-11.5 50t-27.5 28h-32v678q0 27-8 65.5t-23 38.5h-133q-16 0-35-42t-19-62V690H289v678q0 20-19 62t-35 42H102q-15 0-23-38.5t-8-65.5V690H39q-16 0-27.5-28T0 612V471q0-23 22-34zm1034 866V746h-170v312q0 24 9 51.5t25 27.5h71q17 0 41-30.5t24-48.5zm-982 0V746h170v312q0 24-9 51.5t-25 27.5h-71q-17 0-41-30.5t-24-48.5z"/>`),
		g.Group(children),
	)
}