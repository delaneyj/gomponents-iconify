package cib

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ServerFault(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M32 24.245v3.036h-6.521v-3.036zm-32-.141h15.339v3.031H0v-3.036zm17.188 0h6.521v3.031h-6.521v-3.036zm8.291-4.807H32v3.031h-6.521zM0 19.156h15.339v3.031H0zm17.188 0h6.521v3.031h-6.521zm8.291-4.463H32v3.031h-6.521zM0 14.552h15.339v3.031H0zm17.188 0h6.521v3.031h-6.521zm8.291-4.807H32v3.036h-6.521zM0 9.609h15.339v3.031H0V9.598zm17.188 0h6.521v3.031h-6.521V9.598zm8.291-4.744H32v3.036h-6.521V4.859zM0 4.719h15.339v3.036H0zm17.188 0h6.521v3.036h-6.521z"/>`),
		g.Group(children),
	)
}