package ps

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Branch(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M299 0q-35 0-60.5 25T213 85t25.5 60.5T299 171t60-25.5T384 85t-25-60t-60-25zm0 128q-18 0-30.5-12.5T256 85q0-17 12.5-29.5T299 43q17 0 29.5 12.5T341 85q0 18-12.5 30.5T299 128zM149 299h-42V169q27-7 45.5-30.5T171 85q0-35-25.5-60T85 0T25 25T0 85q0 30 18 52.5T64 169v177q-27 6-45.5 29.5T0 429q0 35 25 60t60 25t60.5-25t25.5-60q0-30-18-52.5T107 346v-5h42q71 0 121-50t50-120h-43q0 53-37.5 90.5T149 299zM43 85q0-17 12.5-29.5T85 43q18 0 30.5 12.5T128 85q0 18-12.5 30.5T85 128q-17 0-29.5-12.5T43 85zm85 342q0 17-12.5 29.5T85 469q-17 0-29.5-12.5T43 427q0-18 12.5-30.5T85 384q18 0 30.5 12.5T128 427z"/>`),
		g.Group(children),
	)
}