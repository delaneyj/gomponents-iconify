package el

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Female(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1200 1200"),
		g.Raw(`<path fill="currentColor" d="M600 0C387.136 0 214.538 172.598 214.538 385.462c0 187.751 134.31 344.181 312.055 378.479V888.02H361.278v146.74h165.314V1200h146.74v-165.24h165.314V888.02H673.333V763.94c177.78-34.269 312.129-190.702 312.129-378.479C985.462 172.598 812.864 0 600 0zm0 153.278c128.231 0 232.184 103.953 232.184 232.184S728.232 617.647 600 617.647S367.816 513.693 367.816 385.462S471.769 153.278 600 153.278z"/>`),
		g.Group(children),
	)
}