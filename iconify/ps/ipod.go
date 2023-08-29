package ps

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Ipod(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M299 0H85Q50 0 25 25T0 85v342q0 35 25 60t60 25h214q35 0 60-25t25-60V85q0-35-25-60T299 0zm42 427q0 17-12.5 29.5T299 469H85q-17 0-29.5-12.5T43 427V85q0-17 12.5-29.5T85 43h214q17 0 29.5 12.5T341 85v342zM64 256h256V85H64v171zm43-128h170v85H107v-85zm85 149q-35 0-60 25.5T107 363t25 60t60 25t60-25t25-60t-25-60.5t-60-25.5zm0 128q-17 0-30-12.5T149 363q0-18 13-30.5t30-12.5t30 12.5t13 30.5q0 17-13 29.5T192 405z"/>`),
		g.Group(children),
	)
}