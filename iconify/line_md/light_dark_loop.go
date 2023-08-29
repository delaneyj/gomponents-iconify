package line_md

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LightDarkLoop(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<defs><mask id="lineMdLightDarkLoop0"><circle cx="7.5" cy="7.5" r="5.5" fill="#fff"/><circle cx="7.5" cy="7.5" r="5.5"><animate fill="freeze" attributeName="cx" dur="0.4s" values="7.5;11"/><animate fill="freeze" attributeName="r" dur="0.4s" values="5.5;6.5"/></circle></mask><mask id="lineMdLightDarkLoop1"><g fill="#fff"><circle cx="12" cy="9" r="5.5"><animate fill="freeze" attributeName="cy" begin="1s" dur="0.5s" values="9;15"/></circle><g><g fill-opacity="0"><use href="#lineMdLightDarkLoop2" transform="rotate(-125 12 15)"/><use href="#lineMdLightDarkLoop2" transform="rotate(-75 12 15)"/><use href="#lineMdLightDarkLoop2" transform="rotate(-25 12 15)"/><use href="#lineMdLightDarkLoop2" transform="rotate(25 12 15)"/><use href="#lineMdLightDarkLoop2" transform="rotate(75 12 15)"/><set attributeName="fill-opacity" begin="1.5s" to="1"/></g><animateTransform attributeName="transform" dur="5s" repeatCount="indefinite" type="rotate" values="0 12 15;50 12 15"/></g></g><path d="M0 10h26v5h-26z"/><path fill="none" stroke="#fff" stroke-dasharray="26" stroke-dashoffset="26" stroke-linecap="round" stroke-width="2" d="M1 12h22"><animate attributeName="d" dur="6s" repeatCount="indefinite" values="M0 12h22;M2 12h22;M0 12h22"/><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.5s" dur="0.4s" values="26;52"/></path></mask><symbol id="lineMdLightDarkLoop2"><path d="M0 0"><animate fill="freeze" attributeName="d" begin="1.5s" dur="0.4s" values="M11 18h2L12 20z;M10.5 21.5h3L12 24z"/></path></symbol></defs><g fill="currentColor"><rect width="13" height="13" x="1" y="1" mask="url(#lineMdLightDarkLoop0)"/><path d="M-2 11h28v13h-28z" mask="url(#lineMdLightDarkLoop1)" transform="rotate(-45 12 12)"/></g>`),
		g.Group(children),
	)
}