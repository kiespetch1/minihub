package definitions

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"minihub/pkg/metrics"
)

type BookmarkMetrics struct {
	*metrics.BaseMetrics

	BookmarksTotal   *prometheus.GaugeVec
	BookmarksCreated prometheus.Counter
	BookmarksUpdated prometheus.Counter
	BookmarksDeleted prometheus.Counter
	BookmarksFailed  *prometheus.CounterVec

	BookmarksByCategory *prometheus.GaugeVec
	BookmarksByUser     *prometheus.GaugeVec
	BookmarkDuplicates  prometheus.Counter
}

func NewBookmarkMetrics(namespace, serviceName string) *BookmarkMetrics {
	return &BookmarkMetrics{
		BaseMetrics: metrics.NewBaseMetrics(namespace, serviceName),

		BookmarksTotal: promauto.NewGaugeVec(
			prometheus.GaugeOpts{
				Namespace: namespace,
				Subsystem: serviceName,
				Name:      "bookmarks_total",
				Help:      "Total number of bookmarks in the system",
			},
			[]string{"status"}, // active, archived, deleted
		),

		BookmarksCreated: promauto.NewCounter(
			prometheus.CounterOpts{
				Namespace: namespace,
				Subsystem: serviceName,
				Name:      "bookmarks_created_total",
				Help:      "Total number of bookmarks created",
			},
		),

		BookmarksUpdated: promauto.NewCounter(
			prometheus.CounterOpts{
				Namespace: namespace,
				Subsystem: serviceName,
				Name:      "bookmarks_updated_total",
				Help:      "Total number of bookmarks updated",
			},
		),

		BookmarksDeleted: promauto.NewCounter(
			prometheus.CounterOpts{
				Namespace: namespace,
				Subsystem: serviceName,
				Name:      "bookmarks_deleted_total",
				Help:      "Total number of bookmarks deleted",
			},
		),

		BookmarksFailed: promauto.NewCounterVec(
			prometheus.CounterOpts{
				Namespace: namespace,
				Subsystem: serviceName,
				Name:      "bookmarks_failed_total",
				Help:      "Total number of failed bookmark operations",
			},
			[]string{"operation"}, // create, update, delete, fetch
		),

		BookmarksByCategory: promauto.NewGaugeVec(
			prometheus.GaugeOpts{
				Namespace: namespace,
				Subsystem: serviceName,
				Name:      "bookmarks_by_category",
				Help:      "Number of bookmarks per category",
			},
			[]string{"category"},
		),

		BookmarksByUser: promauto.NewGaugeVec(
			prometheus.GaugeOpts{
				Namespace: namespace,
				Subsystem: serviceName,
				Name:      "bookmarks_by_user",
				Help:      "Number of bookmarks per user",
			},
			[]string{"user_id"},
		),

		BookmarkDuplicates: promauto.NewCounter(
			prometheus.CounterOpts{
				Namespace: namespace,
				Subsystem: serviceName,
				Name:      "bookmark_duplicates_total",
				Help:      "Total number of duplicate bookmark attempts",
			},
		),
	}
}

func (m *BookmarkMetrics) Base() *metrics.BaseMetrics {
	return m.BaseMetrics
}
