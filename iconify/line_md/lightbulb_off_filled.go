package line_md

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LightbulbOffFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<mask id="lineMdLightbulbOffFilled0"><g fill="#fff"><path fill-opacity="0" stroke="#fff" stroke-dasharray="46" stroke-dashoffset="46" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 17H9V14.1973C7.2066 13.1599 6 11.2208 6 9C6 5.68629 8.68629 3 12 3C15.3137 3 18 5.68629 18 9C18 11.2208 16.7934 13.1599 15 14.1973V17z"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.4s" values="46;0"/><animate fill="freeze" attributeName="fill-opacity" begin="0.7s" dur="0.5s" values="0;1"/></path><rect width="6" height="0" x="9" y="20" rx="1"><animate fill="freeze" attributeName="height" begin="0.5s" dur="0.2s" values="0;2"/></rect></g><g fill="none" stroke-dasharray="26" stroke-dashoffset="26" stroke-linecap="round" stroke-width="2" transform="rotate(45 13 12)"><path stroke="#000" d="M0 11h24"/><path stroke="#fff" d="M1 13h22"/><animate fill="freeze" attributeName="stroke-dashoffset" begin="1.2s" dur="0.2s" values="26;0"/></g></mask><rect width="24" height="24" fill="currentColor" mask="url(#lineMdLightbulbOffFilled0)"/>`),
		g.Group(children),
	)
}