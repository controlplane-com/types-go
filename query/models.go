/* auto-generated */

package query

import "github.com/controlplane-com/types-go/base"

type TermOp string

const (
	TermOpEqual            TermOp = "="
	TermOpGreaterThan      TermOp = ">"
	TermOpGreaterThanEqual TermOp = ">="
	TermOpLessThan         TermOp = "<"
	TermOpLessThanEqual    TermOp = "<="
	TermOpNotEqual         TermOp = "!="
	TermOpRegex            TermOp = "~"
	TermOpExists           TermOp = "exists"
	TermOpNotExists        TermOp = "!exists"
)

type SpecMatch string

const (
	SpecMatchAll  SpecMatch = "all"
	SpecMatchAny  SpecMatch = "any"
	SpecMatchNone SpecMatch = "none"
)

type QueryContext map[string]any

type QueryFetch string

const (
	QueryFetchLinks QueryFetch = "links"
	QueryFetchItems QueryFetch = "items"
)

type QueryResultKind string

const (
	QueryResultKindQueryresult QueryResultKind = "queryresult"
)

type Term struct {
	Op       TermOp `json:"op,omitempty"`
	Property string `json:"property,omitempty"`
	Rel      string `json:"rel,omitempty"`
	Tag      string `json:"tag,omitempty"`
	Value    any/* TODO: [object Object]*/ `json:"value,omitempty"`
}

type Spec struct {
	Match SpecMatch `json:"match,omitempty"`
	Terms []Term    `json:"terms,omitempty"`
}

type Query struct {
	Kind    base.Kind    `json:"kind,omitempty"`
	Context QueryContext `json:"context,omitempty"`
	Fetch   QueryFetch   `json:"fetch,omitempty"`
	Spec    Spec         `json:"spec,omitempty"`
}

type QueryResult[T any] struct {
	Kind     QueryResultKind `json:"kind,omitempty"`
	ItemKind base.Kind       `json:"itemKind,omitempty"`
	Items    []T             `json:"items,omitempty"`
	Links    []base.Link     `json:"links,omitempty"`
	Query    Query           `json:"query,omitempty"`
}
