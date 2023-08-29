package healthicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LungsNegative(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none"><g clip-path="url(#healthiconsLungsNegative0)"><path fill="currentColor" fill-rule="evenodd" d="m0 0l48-.04l.04 48l-48 .04L0 0Zm16.542 11.991c5.435.005 5.356 6.873 5.315 10.4l-.004.376c.05-.024.099-.05.148-.078c.44-.256 1.003-.8 1.005-2.197L23.017 6h2v.007l-.011 14.487c-.002 1.397.562 1.941 1 2.198c.05.029.099.055.148.078l-.004-.376c-.034-3.527-.102-10.395 5.333-10.39c7.445.005 12.742 26.004 9.549 28.654c-3.193 2.65-9.574 1.053-12.762-2.133c-2.574-2.572-2.346-8.576-2.183-12.887l.029-.768a4.268 4.268 0 0 1-1.119-.451a3.813 3.813 0 0 1-.994-.835c-.297.352-.64.626-.996.833c-.4.233-.794.37-1.12.45l.028.768c.157 4.31.375 10.315-2.204 12.883c-3.193 3.18-9.576 4.767-12.765 2.112c-3.189-2.655 2.15-28.645 9.596-28.639Z" clip-rule="evenodd"/></g><defs><clipPath id="healthiconsLungsNegative0"><path d="M0 0h48v48H0z"/></clipPath></defs></g>`),
		g.Group(children),
	)
}