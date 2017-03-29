package test

import (
	"io/ioutil"
	"testing"
	"time"

	"github.com/l-vitaly/equifax"
	"github.com/l-vitaly/gounit"
)

func TestBatchFraud(t *testing.T) {
	u := gounit.New(t)

	bf := equifax.NewBatchFraud()

	expected, err := bf.ToCSV([]*equifax.NewApplication{
		{
			ApplicationID:              "333000333",
			ApplicationDate:            equifax.Time{time.Date(2013, 3, 5, 10, 0, 0, 0, time.Local)},
			LastName:                   "Бендер",
			FirstName:                  "Остап",
			MiddleName:                 "Ибрагимович",
			Birthday:                   equifax.Date{time.Date(1970, 2, 1, 0, 0, 0, 0, time.Local)},
			Birthplace:                 "Калининград",
			DocType:                    equifax.DocType1,
			DocNo:                      "1111№222333",
			DocPlace:                   "ОУФМС Ленинского района г. Калининград",
			DocDate:                    equifax.Date{time.Date(2000, 2, 1, 0, 0, 0, 0, time.Local)},
			DocCode:                    "770045",
			PastDocType:                equifax.DocType99,
			Sex:                        equifax.SexType1,
			Citizenship:                equifax.CountryTypeRU,
			INN:                        "123456789012",
			PFR:                        "45623790181",
			DriverNo:                   "11PК222333",
			Education:                  equifax.EducationType4,
			Marital:                    equifax.MaritalType0,
			NumChildren:                "0",
			Email:                      "test1@test.ru",
			HomePhone:                  "4951112233",
			MobilePhone:                "9161002030",
			LaCountry:                  equifax.CountryTypeRU,
			LaIndex:                    "354340",
			LaRegion:                   "Краснодарский край",
			LaDistrict:                 "Адлеровский",
			LaCity:                     "Сочи",
			LaSettlement:               "Адлер",
			LaStreet:                   "Ленина",
			LaHouse:                    "1",
			LaBuilding:                 "3",
			LaStructure:                "2",
			LaApartment:                "13",
			LaYears:                    "12",
			LaMonth:                    "5",
			LaDate:                     equifax.Date{time.Date(2000, 10, 1, 0, 0, 0, 0, time.Local)},
			RaPhone:                    "0112111990",
			RaCountry:                  equifax.CountryTypeRU,
			RaIndex:                    "333444",
			RaRegion:                   "Краснодарский край",
			RaDistrict:                 "Адлеровский",
			RaCity:                     "Сочи",
			RaSettlement:               "Адлер",
			RaStreet:                   "Мира",
			RaHouse:                    "20",
			RaBuilding:                 "1",
			RaStructure:                "1",
			RaApartment:                "10",
			EmployerName:               "ООО \"Хорошие связи\"",
			EmployerSize:               equifax.EmployerSizeType1,
			BusinessIndustry:           equifax.BusinessIndustryType25,
			Position:                   "Главный консультант",
			EmploymentYear:             "5",
			EmploymentMonth:            "5",
			EmploymentDate:             equifax.Date{time.Date(2005, 10, 1, 0, 0, 0, 0, time.Local)},
			EmploymentINN:              "7712345678",
			IncomeProof:                equifax.IncomeProofType1,
			MonthlyIncome:              "32418",
			BaPhone:                    "4957774422",
			BaCountry:                  equifax.CountryTypeRU,
			BaIndex:                    "354340",
			BaRegion:                   "Краснодарский край",
			BaDistrict:                 "Адлеровский",
			BaCity:                     "Сочи",
			BaSettlement:               "Адлер",
			BaStreet:                   "Лескова",
			BaHouse:                    "5",
			BaBuilding:                 "1",
			BaStructure:                "1",
			BaApartment:                "220",
			ProductType:                equifax.ProductTypeType4,
			ProductName:                "Кредитная карта",
			OriginalChannel:            equifax.OriginalChannelType1,
			ProductSumLimit:            "0",
			ProductSumCurrency:         equifax.SumCurrencyTypeRUR,
			DownPaymentAmount:          "0",
			CollateralExistence:        equifax.CollateralExistenceType0,
			CollateralValue:            "0",
			PurchaseExistence:          equifax.PurchaseExistenceType1,
			PurchaseValue:              "0",
			PurchaseModel:              "Корова",
			OperatorCode:               "3468-9790-990-23-RTY",
			OperatorName:               "Иванова Раиса Петровна",
			PosCode:                    "122-210",
			PosPhone:                   "4951234567",
			PosCountry:                 equifax.CountryTypeRU,
			PosIndex:                   "123456",
			PosRegion:                  "Краснодарский край",
			PosDistrict:                "Адлеровский",
			PosCity:                    "Сочи",
			PosSettlement:              "Адлер",
			PosStreet:                  "Ленина",
			PosHouse:                   "25",
			PosBuilding:                "1",
			PosStructure:               "1",
			PosApartment:               "1",
			NewApplicant:               equifax.NewApplicantType1,
			ApplicantType:              equifax.ApplicantTypeType1,
			ApplicantTypeNum:           1,
			ResponseIsNeeded:           equifax.ResponseIsNeededType1,
			ApplicationStatus:          equifax.ApplicationStatusType9,
			ApplicantID:                "1234567891",
			InitialSumLimit:            "0",
			InitialSumCurrency:         equifax.SumCurrencyTypeRUR,
			ApplicationFraudStatus:     equifax.ApplicationFraudStatusType9,
			ApplicationFraudStatusDesc: "Подозрение в мошеничестве",
			DefaultStatus:              equifax.DefaultStatusType9,
		},
	})
	u.AssertNotError(err, "To CSV Err")

	err = ioutil.WriteFile("./batch_expected.csv", []byte(expected), 0755)
	u.AssertNotError(err, "Create expected CSV file")

	u.AssertFileEquals("./batch_expected.csv", "batch_actual.csv", "")
}
