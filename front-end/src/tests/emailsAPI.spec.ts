import { emailsSearch, SearchError } from '@/services/emailsAPI';

describe('emailsAPI module', () => {
  test('exposes emailsSearch as a callable', () => {
    expect(typeof emailsSearch).toBe('function');
  });

  test('SearchError carries message and optional status', () => {
    const err = new SearchError('boom', 500);
    expect(err).toBeInstanceOf(SearchError);
    expect(err.message).toBe('boom');
    expect(err.status).toBe(500);
  });
});
