package cif

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Tv(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 301 201"),
		g.Raw(`<defs><path id="cifTv0" d="M.5.5h150v75H.5z"/><path id="cifTv1" d="M75.5 38h75v37.5L75.5 38zm0 0v37.5H.5l75-37.5zm0 0H.5V.5l75 37.5zm0 0V.5h75L75.5 38z"/></defs><g fill="none" fill-rule="evenodd"><path fill="#5B97B1" fill-rule="nonzero" d="M.5.5h300v150H.5z"/><path fill="#00247D" fill-rule="nonzero" d="M.5.5h150v75H.5z"/><mask id="cifTv2" fill="#fff"><use href="#cifTv0"/></mask><path stroke="#FFF" stroke-width="18" d="m.5.5l150 75m0-75l-150 75" mask="url(#cifTv2)"/><mask id="cifTv3" fill="#fff"><use href="#cifTv1"/></mask><path stroke="#CF142B" stroke-width="12" d="m.5.5l150 75m0-75l-150 75" mask="url(#cifTv3)"/><path stroke="#FFF" stroke-width="30" d="M75.5.5v75M.5 38h150"/><path stroke="#CF142B" stroke-width="18" d="M75.5.5v75M.5 38h150"/><path fill="#FFCE00" fill-rule="nonzero" d="m273.409 18.51l2.8-8.659l2.799 8.659l9.084-.006l-7.354 5.343l2.813 8.654l-7.343-5.356l-7.345 5.355l2.815-8.654l-7.352-5.344zM246.81 94.15l2.8-8.659l2.8 8.659l9.083-.007l-7.353 5.344l2.813 8.654l-7.344-5.356l-7.345 5.355l2.815-8.654l-7.352-5.345zm26.599-17.123l2.8-8.659l2.799 8.659l9.084-.006l-7.354 5.343l2.813 8.654l-7.343-5.356l-7.345 5.355l2.815-8.654l-7.352-5.345zm-114.375 54.86l2.801-8.659l2.799 8.659l9.083-.007l-7.353 5.344l2.813 8.654l-7.344-5.356l-7.344 5.355l2.815-8.654l-7.353-5.345zm41.273-31.892l-2.8 8.658l-2.8-8.659l-9.083.007l7.353-5.343l-2.813-8.655l7.344 5.357l7.344-5.356l-2.814 8.654l7.352 5.345zm31.501-43.566l-2.8 8.659l-2.799-8.659l-9.084.006l7.354-5.343l-2.813-8.654l7.343 5.356l7.345-5.355l-2.815 8.654l7.352 5.344zm26.475-9.048l-2.801 8.658l-2.799-8.658l-9.083.006l7.353-5.343l-2.813-8.655l7.344 5.357l7.344-5.355l-2.815 8.653l7.353 5.345zm-26.475 75.402l-2.8 8.658l-2.799-8.659l-9.084.007l7.354-5.343l-2.813-8.655l7.343 5.357l7.345-5.356l-2.815 8.654l7.352 5.345zm-31.501 7.708l-2.8 8.658l-2.8-8.659l-9.083.007l7.353-5.344l-2.813-8.654l7.344 5.356l7.344-5.355l-2.814 8.654l7.352 5.345z"/></g>`),
		g.Group(children),
	)
}