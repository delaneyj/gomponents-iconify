package line_md

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WeatherCloudyLoop(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<mask id="lineMdWeatherCloudyLoop0"><g fill="#fff"><circle cx="12" cy="11" r="6"><animate attributeName="cx" dur="30s" repeatCount="indefinite" values="12;11;12;13;12"/></circle><rect width="10" height="7" x="8" y="12"/><rect width="16" height="10" x="1" y="9" rx="5"><animate attributeName="x" dur="19s" repeatCount="indefinite" values="1;0;1;2;1"/></rect><rect width="17" height="8" x="6" y="11" rx="4"><animate attributeName="x" dur="23s" repeatCount="indefinite" values="6;5;6;7;6"/></rect></g><circle cx="12" cy="11" r="4"><animate attributeName="cx" dur="30s" repeatCount="indefinite" values="12;11;12;13;12"/></circle><rect width="8" height="6" x="8" y="11"><animate attributeName="x" dur="30s" repeatCount="indefinite" values="8;7;8;9;8"/></rect><rect width="12" height="6" x="3" y="11" rx="3"><animate attributeName="x" dur="19s" repeatCount="indefinite" values="3;2;3;4;3"/></rect><rect width="13" height="4" x="8" y="13" rx="2"><animate attributeName="x" dur="23s" repeatCount="indefinite" values="8;7;8;9;8"/></rect></mask><rect width="24" height="24" fill="currentColor" mask="url(#lineMdWeatherCloudyLoop0)"/>`),
		g.Group(children),
	)
}