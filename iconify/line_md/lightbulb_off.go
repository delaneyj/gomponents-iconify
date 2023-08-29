package line_md

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LightbulbOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<mask id="lineMdLightbulbOff0"><path fill="none" stroke="#fff" stroke-dasharray="46" stroke-dashoffset="46" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 17H9V14.1973C7.2066 13.1599 6 11.2208 6 9C6 5.68629 8.68629 3 12 3C15.3137 3 18 5.68629 18 9C18 11.2208 16.7934 13.1599 15 14.1973V17z"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.4s" values="46;0"/></path><rect width="6" height="0" x="9" y="20" fill="#fff" rx="1"><animate fill="freeze" attributeName="height" begin="0.5s" dur="0.2s" values="0;2"/></rect><g fill="none" stroke-dasharray="26" stroke-dashoffset="26" stroke-linecap="round" stroke-width="2" transform="rotate(45 13 12)"><path stroke="#000" d="M0 11h24"/><path stroke="#fff" d="M1 13h22"/><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.7s" dur="0.2s" values="26;0"/></g></mask><rect width="24" height="24" fill="currentColor" mask="url(#lineMdLightbulbOff0)"/>`),
		g.Group(children),
	)
}