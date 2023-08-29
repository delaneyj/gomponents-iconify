package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Carlcare(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<circle cx="18.627" cy="19.824" r="4.71" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M10.761 24.262a15.37 15.37 0 0 0-.013.628c0 8.297 6.726 15.023 15.024 15.023s15.023-6.726 15.023-15.023"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M17.9 28.091a8.025 8.025 0 0 0 15.731.064M10.758 15.377c-2.44-.85-5.147.457-5.998 2.896s.457 5.148 2.896 5.998a4.712 4.712 0 0 0 3.108-.002"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M23.337 19.824v5.654a2.614 2.614 0 0 0 5.227 0V15.115M43.5 24.244c-1 .239-2.072.25-3.02-.15a5.084 5.084 0 0 1-3.104-4.684v-6.24a5.083 5.083 0 0 1 5.083-5.083M28.564 19.035c0-2.17 2.368-3.918 4.686-3.918"/>`),
		g.Group(children),
	)
}