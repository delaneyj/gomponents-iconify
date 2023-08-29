package flagpack

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Km(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 24"),
		g.Raw(`<defs><path id="flagpackKm0" fill="#fff" d="M0 0h32v24H0z"/></defs><g fill="none"><g clip-path="url(#flagpackKm2)"><use href="#flagpackKm0"/><path fill="#5196ED" fill-rule="evenodd" d="M0 0v24h32V0H0Z" clip-rule="evenodd"/><mask id="flagpackKm1" width="32" height="24" x="0" y="0" maskUnits="userSpaceOnUse" style="mask-type:luminance"><path fill="#fff" fill-rule="evenodd" d="M0 0v24h32V0H0Z" clip-rule="evenodd"/></mask><g fill-rule="evenodd" clip-rule="evenodd" mask="url(#flagpackKm1)"><path fill="#AF0100" d="M0 12v6h32v-6H0Z"/><path fill="#F7FCFF" d="M0 6v6h32V6H0Z"/><path fill="#FECA00" d="M0 0v6h32V0H0Z"/><path fill="#5EAA22" d="m0 0l19 12L0 24V0Z"/><path fill="#F7FCFF" d="M7.305 16.501S3.72 15.278 3.848 11.45c.127-3.83 3.775-4.58 3.775-4.58c-1.277-.923-5.85.159-5.996 4.58c-.146 4.42 4.315 5.402 5.678 5.051Zm.107-6.692l.112-.654l-.475-.464l.657-.096L8 8l.294.595l.657.096l-.475.464l.112.654L8 9.5l-.588.309Zm.112 1.346l-.112.654L8 11.5l.588.309l-.112-.654l.475-.464l-.657-.096L8 10l-.294.595l-.657.096l.475.463Zm-.112 2.654l.112-.654l-.475-.464l.657-.096L8 12l.294.595l.657.096l-.475.463l.112.655L8 13.5l-.588.309Zm0 2l.112-.654l-.475-.464l.657-.096L8 14l.294.595l.657.096l-.475.463l.112.655L8 15.5l-.588.309Z"/></g></g><defs><clipPath id="flagpackKm2"><use href="#flagpackKm0"/></clipPath></defs></g>`),
		g.Group(children),
	)
}