package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Train(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M576 1536q-26 0-45-19t-19-45q0-26 19-45t45-19q26 0 45 19t19 45q0 26-19 45t-45 19zm896 0q-26 0-45-19t-19-45q0-26 19-45t45-19q26 0 45 19t19 45q0 26-19 45t-45 19zM1280 384H768V256h512v128zm165 1408H603l-256 256H165l258-258q-35-4-65-21t-53-42t-36-58t-13-69V192q0-40 15-75t41-61t61-41t75-15h1152q40 0 75 15t61 41t41 61t15 75v1408q0 36-13 68t-35 58t-53 43t-66 21l258 258h-182l-256-256zm219-1152H384v512h1280V640zM448 128q-26 0-45 19t-19 45v320h1280V192q0-26-19-45t-45-19H448zm-64 1472q0 26 19 45t45 19h1152q26 0 45-19t19-45v-320H384v320z"/>`),
		g.Group(children),
	)
}