package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Promplus(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<circle cx="33.793" cy="9.585" r="4.085" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M41.825 27.819c.045-1.012-.997-1.193-1.621-1.028c-1.338.352-1.035 2.804-1.835 3.603c-1.01 1.009-3.022.854-4.576.854s-3.567.155-4.577-.854c-.8-.799-.496-3.251-1.834-3.603c-.624-.165-1.666.016-1.621 1.028c.267 6.074 4.717 5.654 8.032 5.654s7.764.42 8.032-5.654Zm-11.18-12.747h2.113v2.097h-2.113zm4.184 0h2.113v2.097h-2.113zm6.203 10.977c.311-2.08.766-5.758.984-7.783c.08-.747-1.003-2.2-2.48-2.2h-2.219m-10.764 9.983c-.31-2.08-.85-5.751-.983-7.783c-.163-.733 1.002-2.2 2.48-2.2h2.218"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M32.169 25.919v2.73h3.281v-2.73h2.714v-3.281h-2.73v-2.733h-3.282v2.733H29.42v3.28h2.749M8.376 31.795a4.38 4.38 0 0 1 1.983-8.284h0a4.38 4.38 0 0 1 4.38 4.38h0a4.38 4.38 0 0 1-2.397 3.904"/><circle cx="10.359" cy="27.868" r="1.996" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M33.31 33.508C32.077 38.68 27.508 42.5 22.03 42.5c-6.432 0-11.61-5.263-11.61-11.8v-.777"/>`),
		g.Group(children),
	)
}