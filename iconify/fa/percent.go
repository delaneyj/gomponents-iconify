package fa

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Percent(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1536 1536"),
		g.Raw(`<path fill="currentColor" d="M1280 1152q0-52-38-90t-90-38t-90 38t-38 90t38 90t90 38t90-38t38-90zM512 384q0-52-38-90t-90-38t-90 38t-38 90t38 90t90 38t90-38t38-90zm1024 768q0 159-112.5 271.5T1152 1536t-271.5-112.5T768 1152t112.5-271.5T1152 768t271.5 112.5T1536 1152zM1440 64q0 20-13 38L371 1510q-19 26-51 26H160q-26 0-45-19t-19-45q0-20 13-38L1165 26q19-26 51-26h160q26 0 45 19t19 45zM768 384q0 159-112.5 271.5T384 768T112.5 655.5T0 384t112.5-271.5T384 0t271.5 112.5T768 384z"/>`),
		g.Group(children),
	)
}