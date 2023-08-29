package radix_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CodesandboxLogo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M7.711.797a.5.5 0 0 0-.422 0l-6 2.8A.5.5 0 0 0 1 4.05v6.9a.5.5 0 0 0 .289.453l6 2.8a.5.5 0 0 0 .422 0l6-2.8A.5.5 0 0 0 14 10.95v-6.9a.5.5 0 0 0-.289-.453l-6-2.8ZM7.5 3.157L5.98 2.51L7.5 1.8l1.52.71l-1.52.646Zm.196 1.003l2.542-1.08l2.034.949L7.5 6.057L2.728 4.029l2.034-.95L7.304 4.16a.5.5 0 0 0 .392 0ZM8 6.93l5-2.124V7.93l-1.918.882a1 1 0 0 0-.582.908v2.078L8 12.965V6.93Zm3.5 4.402l1.5-.7V9.03l-1.5.69v1.612ZM7 6.93v6.034l-2.498-1.166V9.72a1 1 0 0 0-.582-.908L2 7.929V4.806L7 6.93Zm-5 3.7l1.502.702V9.72L2 9.03v1.602Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}