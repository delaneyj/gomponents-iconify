package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TrainSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M1882 2048h-181l-256-256H602l-256 256H165l259-258q-36-4-67-21t-53-42t-35-58t-13-69V192q0-40 15-75t41-61t61-41t75-15h1152q40 0 75 15t61 41t41 61t15 75v1408q0 36-13 68t-35 58t-53 43t-67 21l258 258zM768 384h512V256H768v128zm896 768V640H384v512h1280zM576 1536q26 0 45-19t19-45q0-26-19-45t-45-19q-26 0-45 19t-19 45q0 26 19 45t45 19zm896 0q26 0 45-19t19-45q0-26-19-45t-45-19q-26 0-45 19t-19 45q0 26 19 45t45 19z"/>`),
		g.Group(children),
	)
}