package flag

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PaOneXOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<defs><clipPath id="flagPa1x10"><path fill-opacity=".7" d="M0 0h512v512H0z"/></clipPath></defs><g fill-rule="evenodd" clip-path="url(#flagPa1x10)"><path fill="#fff" d="M-26-25h592.5v596H-26z"/><path fill="#db0000" d="M255.3-20.4h312.1v275.2h-312z"/><path fill="#0000ab" d="M-54.5 254.8h309.9V571H-54.5zM179 181.6l-46.5-29.2l-46.2 29.5l17.2-48l-46.2-29.6l57.1-.4l17.7-47.8l18.1 47.7h57.1l-45.9 30l17.6 47.8z"/><path fill="#d80000" d="m435.2 449l-46.4-29.2l-46.3 29.5l17.2-48l-46.2-29.5l57.2-.4l17.7-47.8l18 47.7h57.2l-46 30l17.6 47.7z"/></g>`),
		g.Group(children),
	)
}