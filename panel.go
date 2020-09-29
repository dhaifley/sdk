package sdk

/*
   Copyright 2016 Alexander I.Grafov <grafov@gmail.com>
   Copyright 2016-2019 The Grafana SDK authors

   Licensed under the Apache License, Version 2.0 (the "License");
   you may not use this file except in compliance with the License.
   You may obtain a copy of the License at

	   http://www.apache.org/licenses/LICENSE-2.0

   Unless required by applicable law or agreed to in writing, software
   distributed under the License is distributed on an "AS IS" BASIS,
   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
   See the License for the specific language governing permissions and
   limitations under the License.

   ॐ तारे तुत्तारे तुरे स्व
*/

import (
	"bytes"
	"encoding/json"
	"errors"
)

// Each panel may be one of these types.
const (
	CustomType panelType = iota
	DashlistType
	GraphType
	TableType
	TextType
	PluginlistType
	AlertlistType
	SinglestatType
	RowType
)

const MixedSource = "-- Mixed --"

type (
	// Panel represents panels of different types defined in Grafana.
	Panel struct {
		CommonPanel
		// Should be initialized only one type of panels.
		// OfType field defines which of types below will be used.
		*GraphPanel
		*TablePanel
		*TextPanel
		*SinglestatPanel
		*DashlistPanel
		*PluginlistPanel
		*RowPanel
		*AlertlistPanel
		*CustomPanel
	}
	panelType   int8
	CommonPanel struct {
		// Default fields.
		GridPos struct {
			H *float32 `json:"h,omitempty"`
			W *float32 `json:"w,omitempty"`
			X *float32 `json:"x,omitempty"`
			Y *float32 `json:"y,omitempty"`
		} `json:"gridPos"`
		ID    uint   `json:"id"`
		Title string `json:"title"`
		Type  string `json:"type"`

		// Optional fields.
		Collapsed        *bool          `json:"collapsed,omitempty"`
		Datasource       *string        `json:"datasource,omitempty"`
		Description      string         `json:"description,omitempty"`
		Editable         *bool          `json:"editable,omitempty"`
		Error            *bool          `json:"error,omitempty"`
		Height           *FloatOrString `json:"height,omitempty"`
		HideTimeOverride *bool          `json:"hideTimeOverride,omitempty"`
		IsNew            *bool          `json:"isNew,omitempty"`
		Links            []Link         `json:"links,omitempty"`
		MinSpan          *float32       `json:"minSpan,omitempty"`  // templating options
		OfType           panelType      `json:"-"`                  // it is required for defining type of the panel
		Renderer         *string        `json:"renderer,omitempty"` // display styles
		Repeat           *string        `json:"repeat,omitempty"`   // templating options
		// RepeatIteration *int64   `json:"repeatIteration,omitempty"`
		RepeatPanelID *uint `json:"repeatPanelId,omitempty"`
		RepeatedByRow bool  `json:"repeatedByRow,omitempty"`
		ScopedVars    map[string]struct {
			Selected bool   `json:"selected"`
			Text     string `json:"text"`
			Value    string `json:"value"`
		} `json:"scopedVars,omitempty"`
		Span        float32 `json:"span,omitempty"`
		Transparent bool    `json:"transparent,omitempty"`
		Alert       *Alert  `json:"alert,omitempty"`
	}
	AlertEvaluator struct {
		Params []float64 `json:"params,omitempty"`
		Type   string    `json:"type,omitempty"`

		catchall   map[string]interface{} `catchall:"json"`
	}
	AlertOperator struct {
		Type string `json:"type,omitempty"`

		catchall   map[string]interface{} `catchall:"json"`
	}
	AlertQuery struct {
		Params []string `json:"params,omitempty"`

		catchall   map[string]interface{} `catchall:"json"`

	}
	AlertReducer struct {
		Params []string `json:"params,omitempty"`
		Type   string   `json:"type,omitempty"`

		catchall   map[string]interface{} `catchall:"json"`

	}
	AlertCondition struct {
		Evaluator AlertEvaluator `json:"evaluator,omitempty"`
		Operator  AlertOperator  `json:"operator,omitempty"`
		Query     AlertQuery     `json:"query,omitempty"`
		Reducer   AlertReducer   `json:"reducer,omitempty"`
		Type      string         `json:"type,omitempty"`

		catchall   map[string]interface{} `catchall:"json"`
	}
	Alert struct {
		Conditions          []AlertCondition    `json:"conditions,omitempty"`
		ExecutionErrorState string              `json:"executionErrorState,omitempty"`
		Frequency           string              `json:"frequency,omitempty"`
		Handler             int                 `json:"handler,omitempty"`
		Name                string              `json:"name,omitempty"`
		NoDataState         string              `json:"noDataState,omitempty"`
		Notifications       []AlertNotification `json:"notifications,omitempty"`
		Message             string              `json:"message,omitempty"`
		For                 string              `json:"for,omitempty"`

		catchall   map[string]interface{} `catchall:"json"`
	}
	GraphPanel struct {
		AliasColors interface{} `json:"aliasColors"` // XXX
		Bars        bool        `json:"bars"`
		DashLength  *uint       `json:"dashLength,omitempty"`
		Dashes      *bool       `json:"dashes,omitempty"`
		Decimals    *uint       `json:"decimals,omitempty"`
		Fill        int         `json:"fill"`
		//		Grid        grid        `json:"grid"` obsoleted in 4.1 by xaxis and yaxis

		Legend          Legend           `json:"legend,omitempty"`
		LeftYAxisLabel  *string          `json:"leftYAxisLabel,omitempty"`
		Lines           bool             `json:"lines"`
		Linewidth       uint             `json:"linewidth"`
		NullPointMode   string           `json:"nullPointMode"`
		Percentage      bool             `json:"percentage"`
		Pointradius     float32          `json:"pointradius"`
		Points          bool             `json:"points"`
		RightYAxisLabel *string          `json:"rightYAxisLabel,omitempty"`
		SeriesOverrides []SeriesOverride `json:"seriesOverrides,omitempty"`
		SpaceLength     *uint            `json:"spaceLength,omitempty"`
		Stack           bool             `json:"stack"`
		SteppedLine     bool             `json:"steppedLine"`
		Targets         []Target         `json:"targets,omitempty"`
		Thresholds      []Threshold      `json:"thresholds,omitempty"`
		TimeFrom        *string          `json:"timeFrom,omitempty"`
		TimeShift       *string          `json:"timeShift,omitempty"`
		Tooltip         Tooltip          `json:"tooltip"`
		XAxis           bool             `json:"x-axis,omitempty"`
		YAxis           bool             `json:"y-axis,omitempty"`
		YFormats        []string         `json:"y_formats,omitempty"`
		Xaxis           Axis             `json:"xaxis"` // was added in Grafana 4.x?
		Yaxes           []Axis           `json:"yaxes"` // was added in Grafana 4.x?

		catchall   map[string]interface{} `catchall:"json"`
	}
	Threshold struct {
		// the alert threshold value, we do not omitempty, since 0 is a valid
		// threshold
		Value float32 `json:"value"`
		// critical, warning, ok, custom
		ColorMode string `json:"colorMode,omitempty"`
		// gt or lt
		Op   string `json:"op,omitempty"`
		Fill bool   `json:"fill"`
		Line bool   `json:"line"`
		// hexadecimal color (e.g. #629e51, only when ColorMode is "custom")
		FillColor string `json:"fillColor,omitempty"`
		// hexadecimal color (e.g. #629e51, only when ColorMode is "custom")
		LineColor string `json:"lineColor,omitempty"`
		// left or right
		Yaxis string `json:"yaxis,omitempty"`

		catchall   map[string]interface{} `catchall:"json"`
	}

	Tooltip struct {
		Shared       bool   `json:"shared"`
		ValueType    string `json:"value_type"`
		MsResolution bool   `json:"msResolution,omitempty"` // was added in Grafana 3.x
		Sort         int    `json:"sort,omitempty"`

		catchall   map[string]interface{} `catchall:"json"`
	}
	TablePanel struct {
		Columns   []Column      `json:"columns"`
		Sort      *Sort         `json:"sort,omitempty"`
		Styles    []ColumnStyle `json:"styles"`
		Transform string        `json:"transform"`
		Targets   []Target      `json:"targets,omitempty"`
		Scroll    bool          `json:"scroll"` // from grafana 3.x

		catchall   map[string]interface{} `catchall:"json"`
	}
	TextPanel struct {
		// Default.
		Content string `json:"content"`
		Mode    string `json:"mode"`

		// Optional.
		PageSize   uint          `json:"pageSize,omitempty"`
		Scroll     bool          `json:"scroll,omitempty"`
		ShowHeader bool          `json:"showHeader,omitempty"`
		Sort       *Sort         `json:"sort,omitempty"`
		Styles     []ColumnStyle `json:"styles,omitempty"`

		catchall   map[string]interface{} `catchall:"json"`
	}
	SinglestatPanel struct {
		Colors          []string    `json:"colors"`
		ColorValue      bool        `json:"colorValue"`
		ColorBackground bool        `json:"colorBackground"`
		ColorPostfix    bool        `json:"colorPostfix"`
		ColorPrefix     bool        `json:"colorPrefix"`
		Decimals        int         `json:"decimals"`
		Format          string      `json:"format"`
		Gauge           Gauge       `json:"gauge,omitempty"`
		MappingType     *uint       `json:"mappingType,omitempty"`
		MappingTypes    []*MapType  `json:"mappingTypes,omitempty"`
		MaxDataPoints   *IntString  `json:"maxDataPoints,omitempty"`
		NullPointMode   string      `json:"nullPointMode"`
		Postfix         *string     `json:"postfix,omitempty"`
		PostfixFontSize *string     `json:"postfixFontSize,omitempty"`
		Prefix          *string     `json:"prefix,omitempty"`
		PrefixFontSize  *string     `json:"prefixFontSize,omitempty"`
		RangeMaps       []*RangeMap `json:"rangeMaps,omitempty"`
		SparkLine       SparkLine   `json:"sparkline,omitempty"`
		Targets         []Target    `json:"targets,omitempty"`
		Thresholds      string      `json:"thresholds"`
		ValueFontSize   string      `json:"valueFontSize"`
		ValueMaps       []ValueMap  `json:"valueMaps"`
		ValueName       string      `json:"valueName"`

		catchall   map[string]interface{} `catchall:"json"`
	}
	DashlistPanel struct {
		Mode  string   `json:"mode"`
		Limit uint     `json:"limit"`
		Query string   `json:"query"`
		Tags  []string `json:"tags"`

		catchall   map[string]interface{} `catchall:"json"`
	}
	PluginlistPanel struct {
		Limit int `json:"limit,omitempty"`

		catchall   map[string]interface{} `catchall:"json"`
	}
	AlertlistPanel struct {
		OnlyAlertsOnDashboard bool     `json:"onlyAlertsOnDashboard"`
		Show                  string   `json:"show"`
		SortOrder             int      `json:"sortOrder"`
		Limit                 int      `json:"limit"`
		StateFilter           []string `json:"stateFilter"`
		NameFilter            string   `json:"nameFilter,omitempty"`
		DashboardTags         []string `json:"dashboardTags,omitempty"`

		catchall   map[string]interface{} `catchall:"json"`
	}
	RowPanel struct {
		Panels []*Panel `json:"panels,omitempty"`

		catchall   map[string]interface{} `catchall:"json"`
	}
	CustomPanel map[string]interface{}
)

// for a graph panel
type (
	Axis struct {
		Decimals int            `json:"decimals,omitempty"`
		Format   string         `json:"format,omitempty"`
		LogBase  int            `json:"logBase,omitempty"`
		Max      *FloatOrString `json:"max,omitempty"`
		Min      *FloatOrString `json:"min,omitempty"`
		Mode     string         `json:"mode,omitempty"`
		Show     bool           `json:"show"`
		Label    string         `json:"label,omitempty"`

		catchall   map[string]interface{} `catchall:"json"`
	}
	SeriesOverride struct {
		Alias         string      `json:"alias"`
		Bars          *bool       `json:"bars,omitempty"`
		Color         *string     `json:"color,omitempty"`
		Dashes        bool        `json:"dashes,omitempty"`
		Fill          *int        `json:"fill,omitempty"`
		FillBelowTo   *string     `json:"fillBelowTo,omitempty"`
		Legend        *bool       `json:"legend,omitempty"`
		Lines         *bool       `json:"lines,omitempty"`
		Linewidth     *int        `json:"linewidth,omitempty"`
		Stack         *BoolString `json:"stack,omitempty"`
		Transform     *string     `json:"transform,omitempty"`
		YAxis         *int        `json:"yaxis,omitempty"`
		ZIndex        *int        `json:"zindex,omitempty"`
		NullPointMode *string     `json:"nullPointMode,omitempty"`

		catchall   map[string]interface{} `catchall:"json"`
	}
	Sort struct {
		Col  int  `json:"col"`
		Desc bool `json:"desc"`

		catchall   map[string]interface{} `catchall:"json"`
	}
	Legend struct {
		AlignAsTable bool   `json:"alignAsTable"`
		Avg          bool   `json:"avg"`
		Current      bool   `json:"current"`
		HideEmpty    bool   `json:"hideEmpty"`
		HideZero     bool   `json:"hideZero"`
		Max          bool   `json:"max"`
		Min          bool   `json:"min"`
		RightSide    bool   `json:"rightSide"`
		Show         bool   `json:"show"`
		SideWidth    *uint  `json:"sideWidth,omitempty"`
		Sort         string `json:"sort,omitempty"`
		SortDesc     bool   `json:"sortDesc,omitempty"`
		Total        bool   `json:"total"`
		Values       bool   `json:"values"`

		catchall   map[string]interface{} `catchall:"json"`
	}
)

// for a table
type (
	Column struct {
		TextType string `json:"text"`
		Value    string `json:"value"`

		catchall   map[string]interface{} `catchall:"json"`
	}
	ColumnStyle struct {
		Alias           *string   `json:"alias"`
		ColorMode       *string   `json:"colorMode,omitempty"`
		Colors          *[]string `json:"colors,omitempty"`
		DateFormat      *string   `json:"dateFormat,omitempty"`
		Decimals        *int      `json:"decimals,omitempty"`
		Link            *bool     `json:"link,omitempty"`
		LinkTargetBlank *bool     `json:"linkTargetBlank,omitempty"`
		LinkTooltip     *string   `json:"linkTooltip,omitempty"`
		LinkURL         *string   `json:"linkUrl,omitempty"`
		Pattern         string    `json:"pattern"`
		Thresholds      *[]string `json:"thresholds,omitempty"`
		Type            string    `json:"type"`
		Unit            *string   `json:"unit,omitempty"`

		catchall   map[string]interface{} `catchall:"json"`
	}
)

// for a singlestat
type (
	ValueMap struct {
		Op       string `json:"op"`
		TextType string `json:"text"`
		Value    string `json:"value"`

		catchall   map[string]interface{} `catchall:"json"`
	}
	Gauge struct {
		MaxValue         float32 `json:"maxValue"`
		MinValue         float32 `json:"minValue"`
		Show             bool    `json:"show"`
		ThresholdLabels  bool    `json:"thresholdLabels"`
		ThresholdMarkers bool    `json:"thresholdMarkers"`

		catchall   map[string]interface{} `catchall:"json"`
	}
	SparkLine struct {
		FillColor *string  `json:"fillColor,omitempty"`
		Full      bool     `json:"full,omitempty"`
		LineColor *string  `json:"lineColor,omitempty"`
		Show      bool     `json:"show,omitempty"`
		YMin      *float64 `json:"ymin,omitempty"`
		YMax      *float64 `json:"ymax,omitempty"`

		catchall   map[string]interface{} `catchall:"json"`
	}
Metric struct {
ID    string `json:"id"`
Field string `json:"field"`
Type  string `json:"type"`

	catchall   map[string]interface{} `catchall:"json"`
}

	SettingsType struct {
		Interval    string `json:"interval,omitempty"`
		MinDocCount int    `json:"min_doc_count"`
		Order       string `json:"order,omitempty"`
		OrderBy     string `json:"orderBy,omitempty"`
		Size        string `json:"size,omitempty"`

		catchall   map[string]interface{} `catchall:"json"`
	}

	BucketAgg struct {
		ID       string `json:"id"`
		Field    string `json:"field"`
		Type     string `json:"type"`
		Settings SettingsType `json:"settings"`

		catchall   map[string]interface{} `catchall:"json"`
		}
	)

// for an any panel
type Target struct {
	RefID      string `json:"refId"`
	Datasource string `json:"datasource,omitempty"`
	Hide       *bool  `json:"hide,omitempty"`

	// For Prometheus
	Expr           string `json:"expr,omitempty"`
	IntervalFactor int    `json:"intervalFactor,omitempty"`
	Interval       string `json:"interval,omitempty"`
	Step           int    `json:"step,omitempty"`
	LegendFormat   string `json:"legendFormat,omitempty"`
	Instant        bool   `json:"instant,omitempty"`
	Format         string `json:"format,omitempty"`

	// For Elasticsearch
	DsType  *string `json:"dsType,omitempty"`
	Metrics []Metric `json:"metrics,omitempty"`
	Query      string      `json:"query,omitempty"`
	Alias      string      `json:"alias,omitempty"`
	RawQuery   *BoolString `json:"rawQuery,omitempty"`
	TimeField  string      `json:"timeField,omitempty"`
	BucketAggs []BucketAgg `json:"bucketAggs,omitempty"`

	// For Graphite
	Target string `json:"target,omitempty"`

	// For CloudWatch
	Namespace  string            `json:"namespace,omitempty"`
	MetricName string            `json:"metricName,omitempty"`
	Statistics []string          `json:"statistics,omitempty"`
	Dimensions map[string]string `json:"dimensions,omitempty"`
	Period     string            `json:"period,omitempty"`
	Region     string            `json:"region,omitempty"`

	// For the Stackdriver data source. Find out more information at
	// https:/grafana.com/docs/grafana/v6.0/features/datasources/stackdriver/
	AlignOptions       []StackdriverAlignOptions `json:"alignOptions,omitempty"`
	AliasBy            string                    `json:"aliasBy,omitempty"`
	MetricType         string                    `json:"metricType,omitempty"`
	MetricKind         string                    `json:"metricKind,omitempty"`
	Filters            []string                  `json:"filters,omitempty"`
	AlignmentPeriod    string                    `json:"alignmentPeriod,omitempty"`
	CrossSeriesReducer string                    `json:"crossSeriesReducer,omitempty"`
	PerSeriesAligner   string                    `json:"perSeriesAligner,omitempty"`
	ValueType          string                    `json:"valueType,omitempty"`
	GroupBys           []string                  `json:"groupBys,omitempty"`

	// For JSON data source.
	// https://grafana.com/grafana/plugins/simpod-json-datasource
	Data string `json:"data,omitempty"`
	Type string `json:"type,omitempty"`

	catchall   map[string]interface{} `catchall:"json"`
}

// StackdriverAlignOptions defines the list of alignment options shown in
// Grafana during query configuration.
type StackdriverAlignOptions struct {
	Expanded bool                     `json:"expanded"`
	Label    string                   `json:"label"`
	Options  []StackdriverAlignOption `json:"options"`

	catchall   map[string]interface{} `catchall:"json"`
}

// StackdriverAlignOption defines a single alignment option shown in Grafana
// during query configuration.
type StackdriverAlignOption struct {
	Label       string   `json:"label"`
	MetricKinds []string `json:"metricKinds"`
	Text        string   `json:"text"`
	Value       string   `json:"value"`
	ValueTypes  []string `json:"valueTypes"`

	catchall   map[string]interface{} `catchall:"json"`
}

type MapType struct {
	Name  *string `json:"name,omitempty"`
	Value *int    `json:"value,omitempty"`

	catchall   map[string]interface{} `catchall:"json"`
}

type RangeMap struct {
	From *string `json:"from,omitempty"`
	Text *string `json:"text,omitempty"`
	To   *string `json:"to,omitempty"`

	catchall   map[string]interface{} `catchall:"json"`
}

func newTrue() *bool {
	b := true
	return &b
}

func newFalse() *bool {
	b := false
	return &b
}

// NewDashlist initializes panel with a dashlist panel.
func NewDashlist(title string) *Panel {
	if title == "" {
		title = "Panel Title"
	}
	render := "flot"
	return &Panel{
		CommonPanel: CommonPanel{
			OfType:   DashlistType,
			Title:    title,
			Type:     "dashlist",
			Renderer: &render,
			IsNew:    newTrue()},
		DashlistPanel: &DashlistPanel{}}
}

// NewGraph initializes panel with a graph panel.
func NewGraph(title string) *Panel {
	if title == "" {
		title = "Panel Title"
	}
	render := "flot"
	return &Panel{
		CommonPanel: CommonPanel{
			OfType:   GraphType,
			Title:    title,
			Type:     "graph",
			Renderer: &render,
			Span:     12,
			IsNew:    newTrue()},
		GraphPanel: &GraphPanel{
			NullPointMode: "connected",
			Pointradius:   5,
			XAxis:         true,
			YAxis:         true,
		}}
}

// NewTable initializes panel with a table panel.
func NewTable(title string) *Panel {
	if title == "" {
		title = "Panel Title"
	}
	render := "flot"
	return &Panel{
		CommonPanel: CommonPanel{
			OfType:   TableType,
			Title:    title,
			Type:     "table",
			Renderer: &render,
			IsNew:    newTrue()},
		TablePanel: &TablePanel{}}
}

// NewText initializes panel with a text panel.
func NewText(title string) *Panel {
	if title == "" {
		title = "Panel Title"
	}
	render := "flot"
	return &Panel{
		CommonPanel: CommonPanel{
			OfType:   TextType,
			Title:    title,
			Type:     "text",
			Renderer: &render,
			IsNew:    newTrue()},
		TextPanel: &TextPanel{}}
}

// NewSinglestat initializes panel with a singlestat panel.
func NewSinglestat(title string) *Panel {
	if title == "" {
		title = "Panel Title"
	}
	render := "flot"
	return &Panel{
		CommonPanel: CommonPanel{
			OfType:   SinglestatType,
			Title:    title,
			Type:     "singlestat",
			Renderer: &render,
			IsNew:    newTrue()},
		SinglestatPanel: &SinglestatPanel{}}
}

// NewPluginlist initializes panel with a singlestat panel.
func NewPluginlist(title string) *Panel {
	if title == "" {
		title = "Panel Title"
	}
	render := "flot"
	return &Panel{
		CommonPanel: CommonPanel{
			OfType:   PluginlistType,
			Title:    title,
			Type:     "pluginlist",
			Renderer: &render,
			IsNew:    newTrue()},
		PluginlistPanel: &PluginlistPanel{}}
}

func NewAlertlist(title string) *Panel {
	if title == "" {
		title = "Panel Title"
	}
	render := "flot"
	return &Panel{
		CommonPanel: CommonPanel{
			OfType:   AlertlistType,
			Title:    title,
			Type:     "alertlist",
			Renderer: &render,
			IsNew:    newTrue()},
		AlertlistPanel: &AlertlistPanel{}}
}

// NewCustom initializes panel with a singlestat panel.
func NewCustom(title string) *Panel {
	if title == "" {
		title = "Panel Title"
	}
	render := "flot"
	return &Panel{
		CommonPanel: CommonPanel{
			OfType:   CustomType,
			Title:    title,
			Type:     "singlestat",
			Renderer: &render,
			IsNew:    newTrue()},
		CustomPanel: &CustomPanel{}}
}

// ResetTargets delete all targets defined for a panel.
func (p *Panel) ResetTargets() {
	switch p.OfType {
	case GraphType:
		p.GraphPanel.Targets = nil
	case SinglestatType:
		p.SinglestatPanel.Targets = nil
	case TableType:
		p.TablePanel.Targets = nil
	}
}

// AddTarget adds a new target as defined in the argument
// but with refId letter incremented. Value of refID from
// the argument will be used only if no target with such
// value already exists.
func (p *Panel) AddTarget(t *Target) {
	switch p.OfType {
	case GraphType:
		p.GraphPanel.Targets = append(p.GraphPanel.Targets, *t)
	case SinglestatType:
		p.SinglestatPanel.Targets = append(p.SinglestatPanel.Targets, *t)
	case TableType:
		p.TablePanel.Targets = append(p.TablePanel.Targets, *t)
	}
	// TODO check for existing refID
}

// SetTarget updates a target if target with such refId exists
// or creates a new one.
func (p *Panel) SetTarget(t *Target) {
	setTarget := func(t *Target, targets *[]Target) {
		for i, target := range *targets {
			if t.RefID == target.RefID {
				(*targets)[i] = *t
				return
			}
		}
		(*targets) = append((*targets), *t)
	}
	switch p.OfType {
	case GraphType:
		setTarget(t, &p.GraphPanel.Targets)
	case SinglestatType:
		setTarget(t, &p.SinglestatPanel.Targets)
	case TableType:
		setTarget(t, &p.TablePanel.Targets)
	}
}

// MapDatasources on all existing targets for the panel.
func (p *Panel) RepeatDatasourcesForEachTarget(dsNames ...string) {
	repeatDS := func(dsNames []string, targets *[]Target) {
		var refID = "A"
		originalTargets := *targets
		cleanedTargets := make([]Target, 0, len(originalTargets)*len(dsNames))
		*targets = cleanedTargets
		for _, target := range originalTargets {
			for _, ds := range dsNames {
				newTarget := target
				newTarget.RefID = refID
				newTarget.Datasource = ds
				refID = incRefID(refID)
				*targets = append(*targets, newTarget)
			}
		}
	}
	switch p.OfType {
	case GraphType:
		repeatDS(dsNames, &p.GraphPanel.Targets)
	case SinglestatType:
		repeatDS(dsNames, &p.SinglestatPanel.Targets)
	case TableType:
		repeatDS(dsNames, &p.TablePanel.Targets)
	}
}

// RepeatTargetsForDatasources repeats all existing targets for a panel
// for all provided in the argument datasources. Existing datasources of
// targets are ignored.
func (p *Panel) RepeatTargetsForDatasources(dsNames ...string) {
	repeatTarget := func(dsNames []string, targets *[]Target) {
		var lastRefID string
		lenTargets := len(*targets)
		for i, name := range dsNames {
			if i < lenTargets {
				(*targets)[i].Datasource = name
				lastRefID = (*targets)[i].RefID
			} else {
				newTarget := (*targets)[i%lenTargets]
				lastRefID = incRefID(lastRefID)
				newTarget.RefID = lastRefID
				newTarget.Datasource = name
				*targets = append(*targets, newTarget)
			}
		}
	}
	switch p.OfType {
	case GraphType:
		repeatTarget(dsNames, &p.GraphPanel.Targets)
	case SinglestatType:
		repeatTarget(dsNames, &p.SinglestatPanel.Targets)
	case TableType:
		repeatTarget(dsNames, &p.TablePanel.Targets)
	}
}

// GetTargets is iterate over all panel targets. It just returns nil if
// no targets defined for panel of concrete type.
func (p *Panel) GetTargets() *[]Target {
	switch p.OfType {
	case GraphType:
		return &p.GraphPanel.Targets
	case SinglestatType:
		return &p.SinglestatPanel.Targets
	case TableType:
		return &p.TablePanel.Targets
	default:
		return nil
	}
}

type probePanel struct {
	CommonPanel
	//	json.RawMessage
}

func (p *Panel) UnmarshalJSON(b []byte) (err error) {
	var probe probePanel
	if err = json.Unmarshal(b, &probe); err == nil {
		p.CommonPanel = probe.CommonPanel
		switch probe.Type {
		case "graph":
			var graph GraphPanel
			p.OfType = GraphType
			if err = json.Unmarshal(b, &graph); err == nil {
				p.GraphPanel = &graph
			}
		case "table":
			var table TablePanel
			p.OfType = TableType
			if err = json.Unmarshal(b, &table); err == nil {
				p.TablePanel = &table
			}
		case "text":
			var text TextPanel
			p.OfType = TextType
			if err = json.Unmarshal(b, &text); err == nil {
				p.TextPanel = &text
			}
		case "singlestat":
			var singlestat SinglestatPanel
			p.OfType = SinglestatType
			if err = json.Unmarshal(b, &singlestat); err == nil {
				p.SinglestatPanel = &singlestat
			}
		case "dashlist":
			var dashlist DashlistPanel
			p.OfType = DashlistType
			if err = json.Unmarshal(b, &dashlist); err == nil {
				p.DashlistPanel = &dashlist
			}
		case "row":
			var row RowPanel
			p.OfType = RowType
			if err = json.Unmarshal(b, &row); err == nil {
				p.RowPanel = &row
			}
		default:
			var custom = make(CustomPanel)
			p.OfType = CustomType
			if err = json.Unmarshal(b, &custom); err == nil {
				p.CustomPanel = &custom
			}
		}
	}
	return
}

func (p *Panel) MarshalJSON() ([]byte, error) {
	switch p.OfType {
	case GraphType:
		var outGraph = struct {
			CommonPanel
			GraphPanel
		}{p.CommonPanel, *p.GraphPanel}
		return json.Marshal(outGraph)
	case TableType:
		var outTable = struct {
			CommonPanel
			TablePanel
		}{p.CommonPanel, *p.TablePanel}
		return json.Marshal(outTable)
	case TextType:
		var outText = struct {
			CommonPanel
			TextPanel
		}{p.CommonPanel, *p.TextPanel}
		return json.Marshal(outText)
	case SinglestatType:
		var outSinglestat = struct {
			CommonPanel
			SinglestatPanel
		}{p.CommonPanel, *p.SinglestatPanel}
		return json.Marshal(outSinglestat)
	case DashlistType:
		var outDashlist = struct {
			CommonPanel
			DashlistPanel
		}{p.CommonPanel, *p.DashlistPanel}
		return json.Marshal(outDashlist)
	case PluginlistType:
		var outPluginlist = struct {
			CommonPanel
			PluginlistPanel
		}{p.CommonPanel, *p.PluginlistPanel}
		return json.Marshal(outPluginlist)
	case AlertlistType:
		var outAlertlist = struct {
			CommonPanel
			AlertlistPanel
		}{p.CommonPanel, *p.AlertlistPanel}
		return json.Marshal(outAlertlist)
	case RowType:
		var outRow = struct {
			CommonPanel
			RowPanel
		}{p.CommonPanel, *p.RowPanel}
		return json.Marshal(outRow)
	case CustomType:
		var outCustom = customPanelOutput{
			p.CommonPanel,
			*p.CustomPanel,
		}
		return json.Marshal(outCustom)
	}
	return nil, errors.New("can't marshal unknown panel type")
}

type customPanelOutput struct {
	CommonPanel
	CustomPanel
}

func (c customPanelOutput) MarshalJSON() ([]byte, error) {
	b, err := json.Marshal(c.CommonPanel)
	if err != nil {
		return b, err
	}
	// Append custom keys to marshalled CommonPanel.
	buf := bytes.NewBuffer(b[:len(b)-1])

	for k, v := range c.CustomPanel {
		buf.WriteString(`,"`)
		buf.WriteString(k)
		buf.WriteString(`":`)
		b, err := json.Marshal(v)
		if err != nil {
			return b, err
		}
		buf.Write(b)
	}
	buf.WriteString("}")
	return buf.Bytes(), nil
}

func incRefID(refID string) string {
	firstLetter := refID[0]
	ordinal := int(firstLetter)
	ordinal++
	return string(rune(ordinal))
}
