package flat_ui

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Shield(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 100 100"),
		g.Raw(`<path fill="#ECF0F1" fill-rule="evenodd" d="M83.005 65.265C71.94 87.008 50 100 50 100S28.06 87.008 16.995 65.265C5.566 42.806 5 15.784 5 15.784L50 0l45 15.784s-.566 27.022-11.995 49.481z" clip-rule="evenodd"/><path fill="#E64C3C" fill-rule="evenodd" d="M49.99 87c-6.456-4.8-17.698-14.432-24.241-27.342c-6.672-13.166-9.255-28.423-10.243-37.38L50 10.129l34.494 12.149c-.987 8.95-3.569 24.21-10.243 37.379C67.741 72.505 56.448 82.184 49.99 87z" clip-rule="evenodd"/><path fill="#fff" fill-rule="evenodd" d="M62.789 46.041c2.944-2.984 2.942-7.819 0-10.803a7.474 7.474 0 0 0-5.457-2.236c-2.757.047-7.304 3.746-7.304 3.746s-4.675-3.751-7.5-3.747a7.471 7.471 0 0 0-5.317 2.237c-2.942 2.984-2.944 7.819 0 10.803L50 59l12.789-12.959z" clip-rule="evenodd"/><path fill-rule="evenodd" d="M50 0L5 15.784s.566 27.022 11.995 49.481C28.06 87.008 50 100 50 100V0z" clip-rule="evenodd" opacity=".08"/>`),
		g.Group(children),
	)
}