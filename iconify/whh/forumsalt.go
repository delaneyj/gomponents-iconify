package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Forumsalt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M928.784 832h-32v192l-192-192h-416q-53 0-82-46l146-146h384q66 0 113-47t47-113V256h32q40 0 68 28t28 68v384q0 40-28 68t-68 28zm-192-256h-416l-192 192V576h-32q-40 0-68-28t-28-68V96q0-40 28-68t68-28h640q40 0 68 28t28 68v384q0 40-28 68t-68 28z"/>`),
		g.Group(children),
	)
}