package definitions

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"minihub/pkg/metrics"
)

type UserMetrics struct {
	*metrics.BaseMetrics

	UsersTotal         *prometheus.GaugeVec
	UsersRegistered    prometheus.Counter
	UsersDeleted       prometheus.Counter
	LoginAttempts      *prometheus.CounterVec
	LoginSuccessful    prometheus.Counter
	LoginFailed        *prometheus.CounterVec
	PasswordResets     prometheus.Counter
	SessionsActive     prometheus.Gauge
	SessionDuration    prometheus.Histogram
	EmailVerifications *prometheus.CounterVec
}

func NewUserMetrics(namespace, serviceName string) *UserMetrics {
	return &UserMetrics{
		BaseMetrics: metrics.NewBaseMetrics(namespace, serviceName),

		UsersTotal: promauto.NewGaugeVec(
			prometheus.GaugeOpts{
				Namespace: namespace,
				Subsystem: serviceName,
				Name:      "users_total",
				Help:      "Total number of users in the system",
			},
			[]string{"status"}, // active, inactive, banned
		),

		UsersRegistered: promauto.NewCounter(
			prometheus.CounterOpts{
				Namespace: namespace,
				Subsystem: serviceName,
				Name:      "users_registered_total",
				Help:      "Total number of registered users",
			},
		),

		UsersDeleted: promauto.NewCounter(
			prometheus.CounterOpts{
				Namespace: namespace,
				Subsystem: serviceName,
				Name:      "users_deleted_total",
				Help:      "Total number of deleted users",
			},
		),

		LoginAttempts: promauto.NewCounterVec(
			prometheus.CounterOpts{
				Namespace: namespace,
				Subsystem: serviceName,
				Name:      "login_attempts_total",
				Help:      "Total number of login attempts",
			},
			[]string{"method"}, // password, oauth, sso
		),

		LoginSuccessful: promauto.NewCounter(
			prometheus.CounterOpts{
				Namespace: namespace,
				Subsystem: serviceName,
				Name:      "login_successful_total",
				Help:      "Total number of successful logins",
			},
		),

		LoginFailed: promauto.NewCounterVec(
			prometheus.CounterOpts{
				Namespace: namespace,
				Subsystem: serviceName,
				Name:      "login_failed_total",
				Help:      "Total number of failed login attempts",
			},
			[]string{"reason"}, // invalid_password, user_not_found, account_locked
		),

		PasswordResets: promauto.NewCounter(
			prometheus.CounterOpts{
				Namespace: namespace,
				Subsystem: serviceName,
				Name:      "password_resets_total",
				Help:      "Total number of password reset requests",
			},
		),

		SessionsActive: promauto.NewGauge(
			prometheus.GaugeOpts{
				Namespace: namespace,
				Subsystem: serviceName,
				Name:      "sessions_active",
				Help:      "Number of currently active user sessions",
			},
		),

		SessionDuration: promauto.NewHistogram(
			prometheus.HistogramOpts{
				Namespace: namespace,
				Subsystem: serviceName,
				Name:      "session_duration_seconds",
				Help:      "Duration of user sessions in seconds",
				Buckets:   prometheus.ExponentialBuckets(60, 2, 10),
			},
		),

		EmailVerifications: promauto.NewCounterVec(
			prometheus.CounterOpts{
				Namespace: namespace,
				Subsystem: serviceName,
				Name:      "email_verifications_total",
				Help:      "Total number of email verification attempts",
			},
			[]string{"status"}, // sent, verified, failed
		),
	}
}

func (m *UserMetrics) Base() *metrics.BaseMetrics {
	return m.BaseMetrics
}
