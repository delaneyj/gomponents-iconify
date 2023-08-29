package si_glyph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Calculator(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 17 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M15.689.041h-5.606v6.876h6.875V1.299c0-.694-.567-1.258-1.269-1.258zm-.645 4.004h-3.106V2.938h3.106v1.107zM10 15.958h5.674c.71 0 1.284-.569 1.284-1.274V9.042H10v6.916zm1.953-4.989h3.07v1.062h-3.07v-1.062zm0 1.984H15v1.094h-3.047v-1.094zM7.898.041H2.326A1.25 1.25 0 0 0 1.083 1.3v5.623h6.815V.041zM6 4.023h-.984v1.039H3.98V4.023H2.959V2.982H3.98V1.98h1.036v1.002H6v1.041zM1.083 14.676c0 .709.562 1.282 1.255 1.282h5.62V9H1.083v5.676zm2.36-3.967l1.046 1.047l1.048-1.049l.732.732l-1.048 1.049l1.061 1.061l-.735.736l-1.062-1.061l-1.022 1.021l-.732-.732l1.021-1.021l-1.045-1.047l.736-.736z"/>`),
		g.Group(children),
	)
}