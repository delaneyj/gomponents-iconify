package zmdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GlobeLock(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 432 384"),
		g.Raw(`<path fill="currentColor" d="M448 69q9 0 15 6.5t6 15.5v85q0 9-6 15t-15 6H341q-8 0-14.5-6t-6.5-15V91q0-9 6.5-15.5T341 69V59q0-22 16-38t38-16t37.5 16T448 59v10zm-17 0V59q0-15-10.5-26T395 22t-26 11t-11 26v10h73zm-49 171h44q1 12 1 21q0 89-62.5 151.5t-151 62.5t-151-62.5T0 261.5t62.5-151T213 48q33 0 64 10v54q0 18-12.5 30.5T235 155h-43v42q0 9-6.5 15.5T171 219h-43v42h128q9 0 15 6.5t6 15.5v64h22q14 0 25 8t15 21q45-49 45-115q0-7-2-21zM192 431v-42q-18 0-30.5-12.5T149 347v-22L47 223q-4 20-4 38q0 65 42.5 113.5T192 431z"/>`),
		g.Group(children),
	)
}