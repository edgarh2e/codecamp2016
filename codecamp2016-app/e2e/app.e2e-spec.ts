import { Codecamp2016AppPage } from './app.po';

describe('codecamp2016-app App', function() {
  let page: Codecamp2016AppPage;

  beforeEach(() => {
    page = new Codecamp2016AppPage();
  });

  it('should display message saying app works', () => {
    page.navigateTo();
    expect(page.getParagraphText()).toEqual('app works!');
  });
});
