package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PasswordField(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M0 256h2048v1408H0V256zm1920 1280V384H128v1152h1792zM512 1120q-33 0-62-12t-51-35t-34-51t-13-62q0-33 12-62t35-51t51-34t62-13q33 0 62 12t51 35t34 51t13 62q0 33-12 62t-35 51t-51 34t-62 13zm512 0q-33 0-62-12t-51-35t-34-51t-13-62q0-33 12-62t35-51t51-34t62-13q33 0 62 12t51 35t34 51t13 62q0 33-12 62t-35 51t-51 34t-62 13zm512 0q-33 0-62-12t-51-35t-34-51t-13-62q0-33 12-62t35-51t51-34t62-13q33 0 62 12t51 35t34 51t13 62q0 33-12 62t-35 51t-51 34t-62 13z"/>`),
		g.Group(children),
	)
}