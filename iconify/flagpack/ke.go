package flagpack

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Ke(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 24"),
		g.Raw(`<defs><path id="flagpackKe0" fill="#fff" d="M0 0h32v24H0z"/></defs><g fill="none"><g clip-path="url(#flagpackKe1)"><use href="#flagpackKe0"/><path fill="#0067C3" fill-rule="evenodd" d="M0 0h32v24H0V0Z" clip-rule="evenodd"/><path fill="#272727" fill-rule="evenodd" d="M0 0v8h32V0H0Z" clip-rule="evenodd"/><path fill="#4E8B1D" fill-rule="evenodd" d="M0 16v8h32v-8H0Z" clip-rule="evenodd"/><path fill="#E31D1C" stroke="#F7FCFF" stroke-width="2" d="M0 7h-1v10h34V7H0Z"/><ellipse cx="16" cy="12" fill="#AC2317" rx="4" ry="8"/><path fill="#fff" d="M19.623 4.03c.229-.41.843-.859 1.842-1.345a.136.136 0 0 1 .111-.003c.06.026.085.09.056.141L11.274 21.452l-.54-.231l9.13-16.42c-.413-.06-.493-.316-.24-.77Z"/><path fill="#fff" d="M12.366 4.03c.252.455.171.711-.242.77l9.13 16.42l-.54.232L10.358 2.823c-.03-.052-.004-.115.056-.14a.135.135 0 0 1 .11.002c1 .486 1.614.935 1.843 1.345Z"/><path fill="#000" fill-rule="evenodd" d="M19.385 16.265c.39-1.234.615-2.697.615-4.265c0-1.568-.226-3.03-.615-4.265C18.508 8.969 18 10.432 18 12c0 1.568.508 3.03 1.385 4.265Zm-6.77-8.53C12.225 8.969 12 10.432 12 12c0 1.568.226 3.03.615 4.265C13.492 15.031 14 13.569 14 12c0-1.568-.508-3.03-1.385-4.265Z" clip-rule="evenodd"/><path fill="#F7FCFF" fill-rule="evenodd" d="M16.24 11.4c.548-.341.96-1.828.96-3.61c0-1.783-.412-3.27-.96-3.611V11.4Zm-.6-.095c-.487-.47-.84-1.866-.84-3.516s.353-3.046.84-3.515v7.031Zm0 1.495v7.031c-.487-.469-.84-1.866-.84-3.515c0-1.65.353-3.046.84-3.516Zm.358 7.2h.004h-.003Zm.242-.074v-7.22c.548.34.96 1.828.96 3.61c0 1.782-.412 3.269-.96 3.61Z" clip-rule="evenodd"/><path fill="#F7FCFF" d="M16 11.474c.663 0 1.2.471 1.2 1.052c0 .582-.537 1.053-1.2 1.053c-.663 0-1.2-.471-1.2-1.053c0-.581.537-1.052 1.2-1.052Z"/></g><defs><clipPath id="flagpackKe1"><use href="#flagpackKe0"/></clipPath></defs></g>`),
		g.Group(children),
	)
}