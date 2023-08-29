package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Chipotle(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m20.38 12.319l-5.71 1.377l-3.751 3.714s-.644 9.094-.303 11.596s6.973 12.164 8.337 13.263s11.217 1.554 12.695.985s3.372-3.6 3.372-3.6s-5.76-3.07-6.518-4.055s-1.78-5.57-1.78-5.57s4.698-7.01 4.774-8.868s-4.37-7.455-5.305-8.147"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M36.347 10.21c.796-1.024 1.137-2.085 1.137-2.085s-1.175-2.842-1.857-3.335s-4.7-.265-5.268 0s-3.676 3.07-3.676 3.07l-.492 5.154c-.96 1.907-1.642 3.297-3.284 2.943s-2.527-3.638-2.527-3.638"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m15.656 10.702l1.493-.393l2.271 1.202l.96.808l6.303-4.459m3.714 5.154c-1.402.379-2.576.857-3.524.637m-8.312 22.416c-2.172-2.097-3.865-6.973-4.067-7.73s.834-9.247.834-9.247l5.052-.758l4.497.48l2.678 2.4c-1.566 1.87-4.32 7.352-4.32 7.352s1.137 7.932 3.183 10.964c-3.031-1.01-3.991-2.046-4.572-2.981s-4.017-7.68-4.169-8.085s.076-5.659.581-6.265s3.158-.53 3.992.505s-.734 5.407-1.365 6.114s-.884-1.49-.656-2.653"/>`),
		g.Group(children),
	)
}