package helper

import (
	"crud_service/pkg/logger"
	"database/sql"
	"strconv"
	"strings"

	"github.com/golang/protobuf/ptypes/wrappers"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func ReplaceQueryParams(namedQuery string, params map[string]interface{}) (string, []interface{}) {
	var (
		i    int = 1
		args []interface{}
	)

	for k, v := range params {
		if k != "" {
			oldsize := len(namedQuery)
			namedQuery = strings.ReplaceAll(namedQuery, ":"+k, "$"+strconv.Itoa(i))

			if oldsize != len(namedQuery) {
				args = append(args, v)
				i++
			}
		}
	}

	return namedQuery, args
}

func ReplaceSQL(old, searchPattern string) string {
	tmpCount := strings.Count(old, searchPattern)
	for m := 1; m <= tmpCount; m++ {
		old = strings.Replace(old, searchPattern, "$"+strconv.Itoa(m), 1)
	}
	return old
}

func HandleError(log logger.Logger, err error, message string, req interface{}, code codes.Code) error {
	if code != 0 {
		log.Error(message, logger.Error(err), logger.Any("req", req))
		return status.Error(code, message)
	} else if err == sql.ErrNoRows {
		log.Error(message+", Not Found", logger.Error(err), logger.Any("req", req))
		return status.Error(codes.NotFound, "Not Found")
	} else if err != nil {
		log.Error(message, logger.Error(err), logger.Any("req", req))
		return status.Error(codes.Internal, message)
	}
	return nil
}

func ToNullString(s *wrappers.StringValue) (res sql.NullString) {
	if s.GetValue() != "" {
		res.String = s.Value
		res.Valid = true
	}
	return res
}

func ToStringValue(s sql.NullString) *wrappers.StringValue {
	if s.Valid {
		return &wrappers.StringValue{Value: s.String}
	}
	return nil
}

func TotalSumma(totalsum *float64, s string) error {
	sum, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return err
	}
	*totalsum += sum
	return nil
}
