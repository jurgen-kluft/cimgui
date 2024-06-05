#include "cunittest/cunittest.h"


#define ArraySize(x) (sizeof(x) / sizeof(x[0]))

UNITTEST_SUITE_BEGIN(test_volk)
{
    UNITTEST_FIXTURE(main)
    {
        UNITTEST_FIXTURE_SETUP() {}
        UNITTEST_FIXTURE_TEARDOWN() {}

        UNITTEST_TEST(init1)
        {}

        UNITTEST_TEST(init2)
        {}
    }
}
UNITTEST_SUITE_END
